package services

import (
	"fmt"
	"lABoratory/lABoratoryAPI/internal/models"
	"lABoratory/lABoratoryAPI/internal/persistence"
	"lABoratory/lABoratoryAPI/internal/utils"
	"math"
	"regexp"
)

type ExperimentService struct {
	repository        persistence.ExperimentRepository
	securityProvider  utils.SecurityProviderI
	assignmentService AssignmentServiceI
}

type ExperimentServiceI interface {
	GetAll(owner *models.User) ([]models.Experiment, error)
	GetOne(experimentId string, owner *models.User) (*models.Experiment, error)
	Create(experiment models.Experiment) error
	Update(experiment models.Experiment, owner *models.User) error
	Delete(experimentId string, owner *models.User) (bool, error)
	DeleteAll(owner *models.User) (bool, error)
}

func NewExperimentService(r persistence.ExperimentRepository, sp utils.SecurityProviderI, as AssignmentServiceI) ExperimentServiceI {
	e := new(ExperimentService)
	e.repository = r
	e.securityProvider = sp
	e.assignmentService = as
	return e
}

func (s *ExperimentService) GetAll(owner *models.User) ([]models.Experiment, error) {
	return s.repository.GetAll(*owner)
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
	experiments, err := s.GetAll(&experiment.Owner)
	if err != nil {
		return err
	}
	if !s.validateExperiment(experiment, experiments) {
		return fmt.Errorf("bad request")
	}
	err = s.repository.Create(experiment)
	if err != nil {
		return err
	}
	return nil
}

func (s *ExperimentService) Update(experiment models.Experiment, owner *models.User) error {
	experiments, err := s.GetAll(&experiment.Owner)
	if err != nil {
		return err
	}
	if !s.validateExperiment(experiment, experiments) {
		return fmt.Errorf("bad request")
	}
	experimentToUpdate, err := s.repository.GetOne(experiment.Id)
	if err != nil {
		return err
	}
	if !validateOwnership(experimentToUpdate, owner) {
		return fmt.Errorf("ownership error")
	}
	err = s.repository.Update(experiment)
	if err != nil {
		return err
	}
	go s.assignmentService.ResetAssignments(experimentToUpdate, &experiment)
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
	go s.assignmentService.DeleteAll(experimentId)
	return wasDeleted, nil
}

func (s *ExperimentService) DeleteAll(owner *models.User) (bool, error) {
	wasDeleted, err := s.repository.DeleteAll(*owner)
	if err != nil {
		return wasDeleted, err
	}
	go s.deleteAllAssignmentsOfOwner(owner)
	return wasDeleted, nil
}

func (s *ExperimentService) validateExperiment(experiment models.Experiment, existingExperiments []models.Experiment) bool {
	if experiment.Name != "" && !containsExperimentName(existingExperiments, experiment.Name) {
		if isDuplicated(experiment.Assignments) {
			return false
		}
		var acc float64 = 0.0
		regex := regexp.MustCompile("^((a)([1-9]+))|(c)$")
		for _, assig := range experiment.Assignments {
			m := regex.FindAllStringSubmatch(assig.AssignmentName, 1)
			if m == nil {
				return false
			}
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

func (s *ExperimentService) deleteAllAssignmentsOfOwner(owner *models.User) error {
	experiments, err := s.GetAll(owner)
	if err != nil {
		return err
	}
	for _, experiment := range experiments {
		_, err := s.assignmentService.DeleteAll(experiment.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func containsExperimentName(experiments []models.Experiment, experimentName string) bool {
	for _, exp := range experiments {
		if exp.Name == experimentName {
			return true
		}
	}
	return false
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
