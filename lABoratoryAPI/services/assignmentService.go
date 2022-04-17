package services

import (
	"fmt"
	"lABoratory/lABoratoryAPI/models"
	"lABoratory/lABoratoryAPI/persistence"
	"lABoratory/lABoratoryAPI/utils"
	"math"
	"strings"
)

type AssignmentService struct {
	experimentRepository persistence.ExperimentRepository
	customerRepository   persistence.CustomerRepository
	securityProvider     utils.SecurityProviderI
}

type AssignmentServiceI interface {
	SetAssignment(key string, newAssigment string) error
	SetAllAssignments(experiment models.Experiment, newAssigment models.Assignment) error
	GetAssignment(key string) (*models.Customer, error)
	DeleteAll(experimentId string) (bool, error)
	GetAssignmentsOfExperiment(experimentId string) ([]models.Customer, error)
}

func NewAssignmentService(er persistence.ExperimentRepository, cr persistence.CustomerRepository, sp utils.SecurityProviderI) AssignmentServiceI {
	as := new(AssignmentService)
	as.experimentRepository = er
	as.customerRepository = cr
	as.securityProvider = sp
	return as
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
		return as.customerRepository.GetOne(key)
	}
}

func (as AssignmentService) createNewAssignment(experiment *models.Experiment) (*models.Customer, error) {
	existingAssignments, err := as.customerRepository.GetAll(experiment.Id)
	if err != nil {
		return nil, err
	}
	newAssignment, err := as.getNewBalancedAssignment(experiment, &existingAssignments)
	if err != nil {
		return nil, err
	}
	id, err := as.customerRepository.Create(*newAssignment)
	if err != nil {
		return nil, err
	}
	newAssignment.Id = id
	return newAssignment, nil
}

// Function that returns the assignment whose absolute error is the largest depending on
// the current percentages and the theoretical ones of the experiment
func (as AssignmentService) getNewBalancedAssignment(experiment *models.Experiment, existingAssignments *[]models.Customer) (*models.Customer, error) {
	// If an assignment has 100%, return it
	for _, assignment := range experiment.Assignments {
		if int(math.Round(assignment.AssignmentValue)) == 100 {
			return &models.Customer{
				ExperimentId:          experiment.Id,
				AssignmentName:        assignment.AssignmentName,
				AssignmentDescription: assignment.AssignmentDescription,
			}, nil
		}
	}

	// Map that will store the number of customers for each assignment
	count := map[string]int{}
	// Map that will store the current percentages of existing assignments
	currentPercentages := map[string]float64{}

	// Initialize the maps with the assignments that the experiment has
	for _, assignment := range experiment.Assignments {
		count[assignment.AssignmentName] = 0
		currentPercentages[assignment.AssignmentName] = 0.0
	}

	// Count the number of existing customers for each assignment
	for _, customer := range *existingAssignments {
		count[customer.AssignmentName]++
	}

	// Calculate the current percantage of the existing customers
	for assignmentName, value := range count {
		if len(*existingAssignments) == 0 {
			currentPercentages[assignmentName] = 100.0 / float64(len(experiment.Assignments))
		} else {
			currentPercentages[assignmentName] = (float64(value) / float64(len(*existingAssignments))) * 100
		}
	}

	// Get the assignment with the largest positive absolute error based on the current percentages
	// and the theoretical ones of the experiment
	var resultAssignmentName string = experiment.GetRandomAssignment().AssignmentName
	for assignmentName, percentage := range currentPercentages {
		current, err := experiment.GetAssignmentByName(assignmentName)
		if err != nil {
			return nil, err
		}
		other, err := experiment.GetAssignmentByName(resultAssignmentName)
		if err != nil {
			return nil, err
		}
		if (current.AssignmentValue - percentage) > (other.AssignmentValue - currentPercentages[resultAssignmentName]) {
			resultAssignmentName = current.AssignmentName
		}
	}

	// Return the assignment with the largest absolute error
	for _, assignment := range experiment.Assignments {
		if assignment.AssignmentName == resultAssignmentName {
			return &models.Customer{
				ExperimentId:          experiment.Id,
				AssignmentName:        assignment.AssignmentName,
				AssignmentDescription: assignment.AssignmentDescription,
			}, nil
		}
	}
	return nil, fmt.Errorf("error creating new assignment")
}

func (as AssignmentService) SetAssignment(key string, newAssigment string) error {
	assignment, err := as.validateAssignment(key, newAssigment)
	if err != nil {
		return err
	}
	return as.customerRepository.SetAssignment(key, *assignment)
}

func (as AssignmentService) SetAllAssignments(experiment models.Experiment, newAssigment models.Assignment) error {
	_, err := experiment.GetAssignmentByName(newAssigment.AssignmentName)
	if err != nil {
		return err
	}
	return as.customerRepository.SetAllAssignments(experiment.Id, newAssigment)
}

func (as AssignmentService) DeleteAll(experimentId string) (bool, error) {
	return as.customerRepository.DeleteAll(experimentId)
}

func (as AssignmentService) validateAssignment(key string, assignmentName string) (*models.Assignment, error) {
	customer, err := as.GetAssignment(key)
	if err != nil {
		return nil, err
	}
	experiment, err := as.experimentRepository.GetOne(customer.ExperimentId)
	if err != nil {
		return nil, err
	}
	assignment, err := experiment.GetAssignmentByName(assignmentName)
	if err != nil {
		return nil, err
	}
	return assignment, nil
}

func (as AssignmentService) GetAssignmentsOfExperiment(experimentId string) ([]models.Customer, error) {
	return as.customerRepository.GetAll(experimentId)
}
