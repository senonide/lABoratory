package persistence

import "lABoratory/lABoratoryAPI/models"

type CustomerRepository interface {
	GetAll(experimentId string) ([]models.Customer, error)
	GetOne(customerId string) (*models.Customer, error)
	Create(customer models.Customer) (string, error)
	SetAssignment(idCustomer string, newAssigment models.Assignment) error
	SetAllAssignments(experimentId string, newAssigment models.Assignment) error
	DeleteAll(experimentId string) (bool, error)
}
