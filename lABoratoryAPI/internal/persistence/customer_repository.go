package persistence

import "lABoratory/lABoratoryAPI/internal/models"

type CustomerRepository interface {
	GetAll(experimentId string) ([]models.Customer, error)
	GetOne(customerKey, experimentId string) (*models.Customer, error)
	Create(customer models.Customer) (string, error)
	SetAssignment(customerKey, experimentId string, newAssigment models.Assignment, override bool) error
	SetAllAssignments(experimentId string, newAssigment models.Assignment) error
	DeleteAll(experimentId string) (bool, error)
	DeleteAssignment(experimentId string, assignmentKey string) (bool, error)
}
