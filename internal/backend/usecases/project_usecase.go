package usecases

import (
    "github.com/google/uuid"
    "github.com/octo-technology/tezos-link/backend/internal/backend/domain/model"
    "github.com/octo-technology/tezos-link/backend/internal/backend/domain/repository"
    "github.com/sirupsen/logrus"
)

type ProjectUsecase struct{
    repo repository.ProjectRepository
}

type ProjectUsecaseInterface interface{
    SaveProject(name string) (*model.Project, error)
    FindProjects() ([]*model.Project, error)
    FindProject(id int64) (*model.Project, error)
}

func NewProjectUsecase(repo repository.ProjectRepository) *ProjectUsecase {
    return &ProjectUsecase{
        repo:repo,
    }
}

func (pu *ProjectUsecase) SaveProject(name string) (*model.Project, error) {
    p, err := pu.repo.Save(name, uuid.New().String())
    if err != nil {
        logrus.Error("Could not save project", name)
        return nil, err
    }

    logrus.Debug("Saved project", name)
    return p, nil
}

func (pu *ProjectUsecase) FindProjects() ([]*model.Project, error) {
    p, err := pu.repo.FindAll()

    if err != nil {
        logrus.Error("Could not find projects", p)
        return nil, err
    }

    logrus.Debug("Found projects", p)
    return p, nil
}

func (pu *ProjectUsecase) FindProject(id int64) (*model.Project, error) {
    p, err := pu.repo.FindById(id)

    if err != nil {
        logrus.Error("Could not find project", p)
        return nil, err
    }

    logrus.Debug("Found project", p)
    return p, nil
}
