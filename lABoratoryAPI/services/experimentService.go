/*
	This class is in charge of implementing the CRUD methods
*/
package services

import (
	"lABoratory/lABoratoryAPI/database"
	"lABoratory/lABoratoryAPI/models"
)

type ExperimentService struct{}
type IExperimentService interface {
	Create(experiment models.Experiment)
	Read()
	ReadOne(experimentId string)
	Update(experiment models.Experiment, experimentId string)
	Delete(experimentId string)
}

func NewExperimentService() *ExperimentService {
	var e *ExperimentService
	return e
}

func (s *ExperimentService) Create(experiment models.Experiment) error {

	err := database.Create(experiment)

	if err != nil {
		return err
	}

	return nil
}

func (s *ExperimentService) Read() ([]models.Experiment, error) {

	experiments, err := database.Read()

	if err != nil {
		return nil, err
	}

	return experiments, nil
}

func (s *ExperimentService) ReadOne(experimentId string) (models.Experiment, error) {

	experiment, err := database.ReadOne(experimentId)

	if err != nil {
		return experiment, err
	}

	return experiment, nil
}

func (s *ExperimentService) Update(experiment models.Experiment, experimentId string) error {

	err := database.Update(experiment, experimentId)

	if err != nil {
		return err
	}

	return nil
}

func (s *ExperimentService) Delete(experimentId string) error {

	err := database.Delete(experimentId)

	if err != nil {
		return err
	}

	return nil
}
