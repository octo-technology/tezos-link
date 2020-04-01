package usecases

import (
	"errors"
	"github.com/google/uuid"
	modelerrors "github.com/octo-technology/tezos-link/backend/pkg/domain/errors"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	pkgrepository "github.com/octo-technology/tezos-link/backend/pkg/domain/repository"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// ProjectUsecase contains the project repository
type ProjectUsecase struct {
	projectRepo pkgrepository.ProjectRepository
	metricsRepo pkgrepository.MetricsRepository
}

// ProjectUsecaseInterface contains all available methods of the project use-case
type ProjectUsecaseInterface interface {
	CreateProject(name string) (*pkgmodel.Project, error)
	FindProjectAndMetrics(uuid string, from time.Time, to time.Time) (*pkgmodel.Project, *pkgmodel.Metrics, error)
}

// NewProjectUsecase returns a new project use-case
func NewProjectUsecase(projectRepo pkgrepository.ProjectRepository, metricsRepo pkgrepository.MetricsRepository) *ProjectUsecase {
	return &ProjectUsecase{
		projectRepo: projectRepo,
		metricsRepo: metricsRepo,
	}
}

// CreateProject create and save a new project
func (pu *ProjectUsecase) CreateProject(name string) (*pkgmodel.Project, error) {
	creationDate := time.Now().UTC()
	if name == "" {
		logrus.Error("empty project name", name)
		return nil, modelerrors.ErrNoProjectName
	}

	p, err := pu.projectRepo.Save(name, uuid.New().String(), creationDate)
	if err != nil {
		logrus.Error("could not save project", name)
		return nil, err
	}

	logrus.Debug("saved project", name)
	return p, nil
}

// FindProjectAndMetrics finds a project and the associated metrics of a given project's uuid
func (pu *ProjectUsecase) FindProjectAndMetrics(uuid string, from time.Time, to time.Time) (*pkgmodel.Project, *pkgmodel.Metrics, error) {
	p, err := pu.projectRepo.FindByUUID(uuid)

	if errors.Is(err, modelerrors.ErrProjectNotFound) {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	count, err := pu.metricsRepo.CountAll(uuid)
	if err != nil {
		return nil, nil, err
	}

	requestsByDay, err := pu.metricsRepo.FindRequestsByDay(uuid, from, to)
	if err != nil {
		return nil, nil, err
	}
	fullRequestByDayArray := buildFullDateRangeFromRequests(from, to, requestsByDay)

	RPCMetrics, err := pu.metricsRepo.CountRPCPathUsage(uuid, from, to)
	if err != nil {
		return nil, nil, err
	}

	lastRequests, err := pu.metricsRepo.FindLastRequests(uuid)
	if err != nil {
		return nil, nil, err
	}

	m := pkgmodel.NewMetrics(count, fullRequestByDayArray, RPCMetrics, lastRequests)
	logrus.Debug("found project and metrics", p, m)
	return p, &m, nil
}

func buildFullDateRangeFromRequests(from time.Time, to time.Time, computedRequests []*pkgmodel.RequestsByDayMetrics) []*pkgmodel.RequestsByDayMetrics {
	oneDay := 24
	numberOfDays := int(to.Sub(from).Hours()) / oneDay
	fullRequestArray := make([]*pkgmodel.RequestsByDayMetrics, 0, numberOfDays)

	oneDayMore := from.AddDate(0, 0, 1)
	dayCursor := &oneDayMore
	for i := 0; i < numberOfDays; i++ {
		if i != 0 {
			a := dayCursor.AddDate(0, 0, 1)
			dayCursor = &a
		}

		var dayFound = false
		dayFound = checkIfDayIsFoundInRequests(computedRequests, dayCursor, dayFound)

		if dayFound == false {
			fullRequestArray = addDayWithEmptyValue(dayCursor, fullRequestArray)
		}
	}

	return append(fullRequestArray, computedRequests...)
}

func addDayWithEmptyValue(dayCursor *time.Time, fullRequestArray []*pkgmodel.RequestsByDayMetrics) []*pkgmodel.RequestsByDayMetrics {
	emptyMetrics := pkgmodel.NewRequestsByDayMetrics(
		strconv.Itoa(dayCursor.Year()),
		strconv.Itoa(int(dayCursor.Month())),
		strconv.Itoa(dayCursor.Day()),
		0)
	fullRequestArray = append(fullRequestArray, emptyMetrics)
	return fullRequestArray
}

func checkIfDayIsFoundInRequests(computedRequests []*pkgmodel.RequestsByDayMetrics, dayCursor *time.Time, dateFound bool) bool {
	for _, r := range computedRequests {
		if r.Year == strconv.Itoa(dayCursor.Year()) &&
			r.Month == strconv.Itoa(int(dayCursor.Month())) &&
			r.Day == strconv.Itoa(dayCursor.Day()) {
			dateFound = true
		}
	}

	return dateFound
}
