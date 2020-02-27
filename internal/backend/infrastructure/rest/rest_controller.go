package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gamegos/jsend"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
	// Used for the doamin objects to be found by Swagger
	_ "github.com/octo-technology/tezos-link/backend/internal/backend/domain/model"
	"github.com/octo-technology/tezos-link/backend/internal/backend/infrastructure/rest/inputs"
	"github.com/octo-technology/tezos-link/backend/internal/backend/usecases"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"strconv"
)

// Controller represents a REST controller
type Controller struct {
	router *chi.Mux
	pu     usecases.ProjectUsecaseInterface
	hu     usecases.HealthUsecaseInterface
}

// NewRestController returns a new rest controller
func NewRestController(router *chi.Mux, pu usecases.ProjectUsecaseInterface, hu usecases.HealthUsecaseInterface) *Controller {
	return &Controller{
		router: router,
		pu:     pu,
		hu:     hu,
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
			r.Get("/", rc.GetProjects)
			r.Post("/", rc.PostProject)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", rc.GetProject)
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
	h := rc.hu.Health()
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
	errDecoding := json.NewDecoder(r.Body).Decode(&inputProject)
	if errDecoding != nil {
		_, _ = jsend.Wrap(w).Data(errDecoding.Error()).Status(http.StatusBadRequest).Send()
		return
	}

	p, errSaving := rc.pu.SaveProject(inputProject.Name)
	if errSaving != nil {
		_, _ = jsend.Wrap(w).Data(errSaving.Error()).Status(http.StatusBadRequest).Send()
		return
	}

	_, _ = jsend.Wrap(w).Data(p).Status(http.StatusCreated).Send()
}

// GetProject godoc
// @Summary Get a Project
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} model.Project
// @Router /projects/{id} [get]
func (rc *Controller) GetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	p, err := rc.pu.FindProject(id)
	if err != nil {
		_, _ = jsend.Wrap(w).Data(err.Error()).Status(http.StatusBadRequest).Send()
		return
	}

	_, _ = jsend.Wrap(w).Data(p).Status(http.StatusOK).Send()
}

// GetProjects godoc
// @Summary Get all Projects
// @Produce json
// @Success 200 {object} model.Project[]
// @Router /projects [get]
func (rc *Controller) GetProjects(w http.ResponseWriter, r *http.Request) {
	p, err := rc.pu.FindProjects()
	if err != nil {
		_, _ = jsend.Wrap(w).Data(err.Error()).Status(http.StatusBadRequest).Send()
		return
	}

	_, _ = jsend.Wrap(w).Data(p).Status(http.StatusOK).Send()
}
