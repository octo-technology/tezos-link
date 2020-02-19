package http

import (
    "fmt"
    "github.com/go-chi/chi"
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
    "regexp"
    "strings"
    "time"
)

type httpController struct {
    router              *chi.Mux
    uc                  usecases.ProxyUsecaseInterface
    reverseProxy        *httputil.ReverseProxy
    httpServer          *http.Server
    UUIDRegexp          *regexp.Regexp
}

const (
    UUIDRegex       = `(?m)([0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12})`
)

func NewHttpController(uc usecases.ProxyUsecaseInterface, rp *httputil.ReverseProxy, srv *http.Server) *httpController {
    return &httpController{
        uc:             uc,
        reverseProxy:   rp,
        httpServer:     srv,
        UUIDRegexp:     regexp.MustCompile(UUIDRegex),
    }
}

func (p *httpController) Initialize() {
    basePath := "v1/"
    middleware := setupLimiterMiddleware()
    http.Handle("/" + basePath, middleware.Handler(http.HandlerFunc(handleProxying(p, basePath))))
}

func (p *httpController) Run() {
    log.Fatal(p.httpServer.ListenAndServe())
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

func handleProxying(p *httpController, basePath string) func(w http.ResponseWriter, req *http.Request) {
    return func(w http.ResponseWriter, req *http.Request) {
        var request = model.NewRequest(
            getRPCFromPath(basePath, req.URL.Path, p.UUIDRegexp),
            getUUIDFromPath(req.URL.Path, p.UUIDRegexp),
            model.MethodFromString(req.Method),
            req.RemoteAddr)
        logrus.Debug(request.Method, request.Path, request.UUID, request.RemoteAddr)

        r, toRawProxy, err := p.uc.Proxy(&request)
        if err != nil {
            logrus.Error(fmt.Sprintf("could not proxy request: %s", err))
        }

        if toRawProxy {
            forwardRawRequestAndRespond(p, w, req)
            return
        }

        respondToRequest(w, r)
    }
}

func getUUIDFromPath(path string, re *regexp.Regexp) string {
    var rpcPath string
    for _, match := range re.FindAllString(path, -1) {
        rpcPath = match
    }
    return rpcPath
}

func getRPCFromPath(basePath string, path string, re *regexp.Regexp) string {
    return strings.Replace(path, basePath + getUUIDFromPath(path, re), "", -1)
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
