package services

import (
	"fmt"
	"lABoratory/lABoratoryAPI/models"
	"lABoratory/lABoratoryAPI/persistence"
	"math"
)

type ExperimentService struct {
	repository persistence.ExperimentRepository
}

type ExperimentServiceI interface {
	GetAll() ([]models.Experiment, error)
	GetOne(experimentId string) (*models.Experiment, error)
	Create(experiment models.Experiment) error
	Update(experiment models.Experiment) error
	Delete(experimentId string) (bool, error)
	DeleteAll(owner *models.User) (bool, error)
}

func NewExperimentService(r persistence.ExperimentRepository) *ExperimentService {
	e := new(ExperimentService)
	e.repository = r
	return e
}

func (s *ExperimentService) GetAll(owner *models.User) ([]models.Experiment, error) {
	experiments, err := s.repository.GetAll(*owner)
	if err != nil {
		return nil, err
	}
	return experiments, nil
}

func (s *ExperimentService) GetOne(experimentId string, owner *models.User) (*models.Experiment, error) {
	experiment, err := s.repository.GetOne(experimentId)
	if err != nil {
		return nil, err
	}
	if validateOwnership(experiment, owner) {
		return experiment, nil
	} else {
		return nil, fmt.Errorf("ownership error")
	}

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

func (s *ExperimentService) Update(experiment models.Experiment, owner *models.User) error {
	if !validateExperiment(experiment) {
		return fmt.Errorf("bad request")
	}
	if !validateOwnership(&experiment, owner) {
		return fmt.Errorf("ownership error")
	}
	err := s.repository.Update(experiment)
	if err != nil {
		return err
	}
	return nil
}

func (s *ExperimentService) Delete(experimentId string, owner *models.User) (bool, error) {
	experiment, err := s.repository.GetOne(experimentId)
	if err != nil {
		return false, err
	}
	if !validateOwnership(experiment, owner) {
		return false, fmt.Errorf("ownership error")
	}
	wasDeleted, err := s.repository.Delete(experimentId)
	if err != nil {
		return wasDeleted, err
	}
	return wasDeleted, nil
}

func (s *ExperimentService) DeleteAll(owner *models.User) (bool, error) {
	wasDeleted, err := s.repository.DeleteAll(*owner)
	if err != nil {
		return wasDeleted, err
	}
	return wasDeleted, nil
}

func validateExperiment(experiment models.Experiment) bool {
	if experiment.Name != "" {
		if isDuplicated(experiment.Assignments) {
			return false
		}
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

func isDuplicated(arr []models.Assignment) bool {
	visited := make(map[models.Assignment]bool, 0)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] {
			return true
		} else {
			visited[arr[i]] = true
		}
	}
	return false
}

func validateOwnership(experiment *models.Experiment, owner *models.User) bool {
	if experiment.Owner == *owner {
		return true
	} else {
		return false
	}
}
