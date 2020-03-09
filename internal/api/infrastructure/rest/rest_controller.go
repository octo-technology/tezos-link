package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gamegos/jsend"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/octo-technology/tezos-link/backend/internal/api/infrastructure/rest/inputs"
	"github.com/octo-technology/tezos-link/backend/internal/api/infrastructure/rest/outputs"
	modelerrors "github.com/octo-technology/tezos-link/backend/pkg/domain/errors"
	"time"

	// Used for the output objects to be found by Swagger
	_ "github.com/octo-technology/tezos-link/backend/internal/api/infrastructure/rest/outputs"
	// Used for the health object to be found by Swagger
	_ "github.com/octo-technology/tezos-link/backend/internal/api/domain/model"
	"github.com/octo-technology/tezos-link/backend/internal/api/usecases"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// Controller represents a REST controller
type Controller struct {
	router         *chi.Mux
	projectUsecase usecases.ProjectUsecaseInterface
	healthUsecase  usecases.HealthUsecaseInterface
}

// NewRestController returns a new rest controller
func NewRestController(
	router *chi.Mux,
	projectUsecase usecases.ProjectUsecaseInterface,
	healthUsecase usecases.HealthUsecaseInterface) *Controller {
	return &Controller{
		router:         router,
		projectUsecase: projectUsecase,
		healthUsecase:  healthUsecase,
	}
}

// @title Tezos Link API
// @version v1
// @description API to manage projects
// @contact.name API Support
// @contact.email email@ded.fr
// @BasePath /api/v1

// Initialize initialize the routes
func (rc *Controller) Initialize() {
	cors := cors.New(cors.Options{
		// TODO Filter to only allow tezoslink.io origin
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		ExposedHeaders: []string{"Location"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	})

	rc.router.Use(cors.Handler)
	rc.router.Use(middleware.RequestID)
	rc.router.Use(middleware.Logger)
	rc.router.Use(middleware.Recoverer)

	// Private routes
	rc.router.Group(func(r chi.Router) {
		// TODO: Add auth
		//tokenAuth := jwtauth.New("HS256", []byte(config.BackendConfig.Jwt.SignKey), nil)
		//r.Use(jwtauth.Verifier(tokenAuth))
		//r.Use(jwtauth.Authenticator)

		r.Route("/api/v1/projects", func(r chi.Router) {
			r.Post("/", rc.PostProject)

			r.Route("/{uuid}", func(r chi.Router) {
				r.Get("/", rc.GetProjectWithMetrics)
				// TODO: To implement
				// r.Put("/", rc.UpdateProject)
				// r.Delete("/", rc.DeleteProject)
			})
		})
	})

	// Public routes
	rc.router.Get("/documentation/", httpSwagger.WrapHandler)
	rc.router.Get("/health", rc.GetHealth)

	_ = chi.Walk(rc.router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Println(method, route)
		return nil
	})
}

// Run runs the controller
func (rc *Controller) Run(port int) {
	logrus.Info("running HTTP API on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), rc.router))
}

// GetHealth godoc
// @Summary get application health
// @Success 200 {object} model.Health
// @Router /health [get]
func (rc *Controller) GetHealth(w http.ResponseWriter, r *http.Request) {
	h := rc.healthUsecase.Health()
	_, _ = jsend.Wrap(w).Data(h).Status(http.StatusOK).Send()
}

// PostProject godoc
// @Summary Create a Project
// @Produce json
// @Param new-project body inputs.NewProject true "New Project"
// @Success 201
// @Failure 400
// @Router /projects [post]
func (rc *Controller) PostProject(w http.ResponseWriter, r *http.Request) {
	var inputProject inputs.NewProject
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	errDecoding := decoder.Decode(&inputProject)
	if errDecoding != nil {
		_, _ = jsend.Wrap(w).Data(errDecoding.Error()).Status(http.StatusBadRequest).Send()
		return
	}

	p, errSaving := rc.projectUsecase.CreateProject(inputProject.Title)
	if errSaving != nil {
		_, _ = jsend.Wrap(w).Data(errSaving.Error()).Status(http.StatusBadRequest).Send()
		return
	}

	w.Header().Add("Location", p.UUID)
	_, _ = jsend.Wrap(w).Status(http.StatusCreated).Send()
}

// GetProjectWithMetrics godoc
// @Summary Get a Project with the associated metrics
// @Produce json
// @Param uuid path string true "Project UUID"
// @Success 200 {object} outputs.ProjectOutputWithMetrics
// @Router /projects/{uuid} [get]
func (rc *Controller) GetProjectWithMetrics(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	now := time.Now()
	nowMinusOneMonth := now.AddDate(0, -1, 0)
	project, metrics, err := rc.projectUsecase.FindProjectAndMetrics(uuid, nowMinusOneMonth, now)
	if errors.Is(err, modelerrors.ErrProjectNotFound) {
		_, _ = jsend.Wrap(w).Data(err.Error()).Status(http.StatusNotFound).Send()
		return
	}
	if err != nil {
		_, _ = jsend.Wrap(w).Data(err.Error()).Status(http.StatusBadRequest).Send()
		return
	}

	po := outputs.NewProjectOutputWithMetrics(project, metrics)
	_, _ = jsend.Wrap(w).Data(po).Status(http.StatusOK).Send()
}
