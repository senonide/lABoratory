/*
	This class is in charge of implementing the CRUD methods
*/
package services

import (
	"fmt"
	"lABoratory/lABoratoryAPI/models"
	"lABoratory/lABoratoryAPI/persistence"
	"lABoratory/lABoratoryAPI/persistence/database"
	"math"
)

type ExperimentService struct {
	repository persistence.ExperimentRepository
}

func NewExperimentService() *ExperimentService {
	e := new(ExperimentService)
	e.repository = database.NewDbExperimentRepository()
	return e
}

func (s *ExperimentService) Create(experiment models.Experiment) error {
	if !validateExperiment(experiment) {
		return fmt.Errorf("bad request")
	}
	err := s.repository.Create(experiment)
	if err != nil {
		return err
	}
	return nil
}

func (s *ExperimentService) GetAll() ([]models.Experiment, error) {
	experiments, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return experiments, nil
}

func (s *ExperimentService) GetOne(experimentId string) (*models.Experiment, error) {
	experiment, err := s.repository.GetOne(experimentId)
	if err != nil {
		return nil, err
	}
	return experiment, nil
}

func (s *ExperimentService) Update(experiment models.Experiment) error {
	if !validateExperiment(experiment) {
		return fmt.Errorf("bad request")
	}
	err := s.repository.Update(experiment)
	if err != nil {
		return err
	}
	return nil
}

func (s *ExperimentService) Delete(experimentId string) (bool, error) {
	wasDeleted, err := s.repository.Delete(experimentId)
	if err != nil {
		return wasDeleted, err
	}
	return wasDeleted, nil
}

func validateExperiment(experiment models.Experiment) bool {
	if experiment.Name != "" {
		var acc float64 = 0.0
		for _, assig := range experiment.Assignments {
			if assig.AssignmentName == "" {
				return false
			}
			acc += assig.AssignmentValue
		}
		totalPercentaje := int(math.Round(acc))
		if totalPercentaje == 100 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
