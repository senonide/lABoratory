package persistence

import "lABoratory/lABoratoryAPI/internal/models"

type ExperimentRepository interface {
	GetAll(owner models.User) ([]models.Experiment, error)
	GetOne(experimentId string) (*models.Experiment, error)
	Create(experiment models.Experiment) error
	Update(experiment models.Experiment) error
	Delete(experimentId string) (bool, error)
	DeleteAll(owner models.User) (bool, error)
}
