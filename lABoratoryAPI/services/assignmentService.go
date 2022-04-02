package services

import (
	"fmt"
	"lABoratory/lABoratoryAPI/models"
	"lABoratory/lABoratoryAPI/persistence"
	"lABoratory/lABoratoryAPI/utils"
	"math/rand"
	"strings"
)

type AssignmentService struct {
	experimentRepository persistence.ExperimentRepository
	customerRepository   persistence.CustomerRepository
	securityProvider     utils.SecurityProviderI
}

type AssignmentServiceI interface {
	SetAssignment(key string, newAssigment *models.Assignment) error
	GetAssignment(key string) (*models.Customer, error)
	Update(updatedExperiment models.Experiment) error
	DeleteAll(experimentId string) (bool, error)
	DeleteAllOfOwner(owner *models.User) (bool, error)
}

func NewAssignmentService(er persistence.ExperimentRepository, cr persistence.CustomerRepository, sp utils.SecurityProviderI) AssignmentServiceI {
	as := new(AssignmentService)
	as.experimentRepository = er
	as.customerRepository = cr
	as.securityProvider = sp
	return as
}

func (as AssignmentService) SetAssignment(key string, newAssigment *models.Assignment) error {
	return as.customerRepository.SetAssignment(key, *newAssigment)
}

func (as AssignmentService) GetAssignment(key string) (*models.Customer, error) {
	if strings.Contains(key, "experimentKey_") {
		token, err := as.securityProvider.GetToken(strings.Replace(key, "experimentKey_", "", 1))
		if err != nil {
			return nil, err
		}
		if !as.securityProvider.ValidateToken(token) {
			return nil, fmt.Errorf("invalid experiment key")
		}
		claims, err := as.securityProvider.GetTokenClaims(token)
		if err != nil {
			return nil, err
		}
		experimentId, ok := claims["sub"].(string)
		if !ok {
			return nil, fmt.Errorf("error decoding token claims")
		}
		experiment, err := as.experimentRepository.GetOne(experimentId)
		if err != nil {
			return nil, err
		}
		return as.createNewAssignment(experiment)
	} else {
		return as.getExistingAssignment(key)
	}
}

func (as AssignmentService) Update(updatedExperiment models.Experiment) error {
	//TODO: Force the update for all assignments
	return nil
}

func (as AssignmentService) DeleteAll(experimentId string) (bool, error) {
	return as.customerRepository.DeleteAll(experimentId)
}

func (as AssignmentService) DeleteAllOfOwner(owner *models.User) (bool, error) {
	//TODO: Get all the experiments of the owner and delete its assignmnets
	return false, nil
}

func (as AssignmentService) createNewAssignment(experiment *models.Experiment) (*models.Customer, error) {
	var targetAssignment models.Assignment
	target := rand.Float64() * 100
	acc := 0.0
	for _, assignment := range experiment.Assignments {
		if target <= (assignment.AssignmentValue + acc) {
			targetAssignment = assignment
			break
		}
		acc += assignment.AssignmentValue
	}
	newAssigment := models.Customer{
		ExperimentId:          experiment.Id,
		AssignmentName:        targetAssignment.AssignmentName,
		AssignmentDescription: targetAssignment.AssignmentDescription,
	}
	id, err := as.customerRepository.Create(newAssigment)
	if err != nil {
		return nil, err
	}
	newAssigment.Id = id
	return &newAssigment, nil
}

func (as AssignmentService) getExistingAssignment(key string) (*models.Customer, error) {
	return as.customerRepository.GetOne(key)
}
