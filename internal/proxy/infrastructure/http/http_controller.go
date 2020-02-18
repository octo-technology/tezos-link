package http

import (
    "fmt"
    "github.com/octo-technology/tezos-link/backend/config"
    "github.com/octo-technology/tezos-link/backend/internal/proxy/domain/model"
    "github.com/octo-technology/tezos-link/backend/internal/proxy/usecases"
    "github.com/sirupsen/logrus"
    "github.com/ulule/limiter"
    "github.com/ulule/limiter/drivers/middleware/stdlib"
    "github.com/ulule/limiter/drivers/store/memory"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "regexp"
    "strconv"
    "time"
)

type httpController struct {
    uc                  usecases.ProxyUsecaseInterface
    reverseProxy        *httputil.ReverseProxy
    UUIDRegexp          *regexp.Regexp
    RPCPathRegexp       *regexp.Regexp
}

const (
    UUIDRegex       = `(?m)([0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12})`
    RPCPathRegexp   = `[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}(.*)`
)

func NewHttpController(uc usecases.ProxyUsecaseInterface) *httpController {
    reverseUrl, err := url.Parse("http://" + config.ProxyConfig.Tezos.Host + ":" + strconv.Itoa(config.ProxyConfig.Tezos.Port))
    if err != nil {
        log.Fatal(fmt.Sprintf("could not read blockchain node reverse url from configuration: %s", err))
    }

    return &httpController{
        uc:             uc,
        reverseProxy:   httputil.NewSingleHostReverseProxy(reverseUrl),
        UUIDRegexp:     regexp.MustCompile(UUIDRegex),
        RPCPathRegexp:  regexp.MustCompile(RPCPathRegexp),
    }
}

func (p *httpController) Run() {
    srv := http.Server{
        Addr:         ":" + strconv.Itoa(config.ProxyConfig.Server.Port),
        ReadTimeout:  time.Duration(config.ProxyConfig.Proxy.ReadTimeout) * time.Second,
        WriteTimeout: time.Duration(config.ProxyConfig.Proxy.WriteTimeout) * time.Second,
        IdleTimeout:  time.Duration(config.ProxyConfig.Proxy.IdleTimeout) * time.Second,
    }
    middleware := setupLimiterMiddleware()
    handlerFunc := handleProxying(p)

    http.Handle("/v1/", middleware.Handler(http.HandlerFunc(handlerFunc)))

    err := srv.ListenAndServe()
    if err != nil {
        log.Fatal(fmt.Sprintf("could not launch proxy: %s", err))
    }
}

func setupLimiterMiddleware() *stdlib.Middleware {
    rate := limiter.Rate{
        Period: time.Duration(config.ProxyConfig.Proxy.RateLimitPeriod) * time.Second,
        Limit:  config.ProxyConfig.Proxy.RateLimitCount,
    }
    store := memory.NewStore()
    middleware := stdlib.NewMiddleware(limiter.New(store, rate), stdlib.WithForwardHeader(true))
    return middleware
}

func handleProxying(p *httpController) func(w http.ResponseWriter, req *http.Request) {
    return func(w http.ResponseWriter, req *http.Request) {
        var request = model.NewRequest(
            getRPCFromPath(req.URL.Path, p.UUIDRegexp),
            getUUIDFromPath(req.URL.Path, p.RPCPathRegexp),
            model.MethodFromString(req.Method),
            req.RemoteAddr)
        logrus.Debug(request.Method, request.Path, request.UUID, request.RemoteAddr)

        r, toRawProxy, err := p.uc.Proxy(&request)
        if err != nil {
            logrus.Error(fmt.Sprintf("could not proxy request: %s", err))
        }

        // TODO: Test logic
        if toRawProxy {
            forwardRawRequestAndRespond(p, w, req)
            return
        }

        respondToRequest(w, r)
    }
}

func getUUIDFromPath(path string, re *regexp.Regexp) string {
    var UUID string
    for _, match := range re.FindAllString(path, -1) {
        UUID = match
    }
    return UUID
}

func getRPCFromPath(path string, re *regexp.Regexp) string {
    var rpcPath string
    for _, match := range re.FindAllString(path, -1) {
        rpcPath = match
    }
    return rpcPath
}

func forwardRawRequestAndRespond(p *httpController, w http.ResponseWriter, req *http.Request) {
    p.reverseProxy.ServeHTTP(w, req)
}

func respondToRequest(w http.ResponseWriter, r string) {
    optionsHeaders(w)
    _, _ = fmt.Fprint(w, r)
}

func optionsHeaders(w http.ResponseWriter) {
    w.Header().Set("Allow", "OPTIONS, POST")
    w.Header().Set("Accept", "application/json")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Content-Type", "application/json")
}
