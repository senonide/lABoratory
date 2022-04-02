package persistence

import "lABoratory/lABoratoryAPI/models"

type CustomerRepository interface {
	GetOne(customerId string) (*models.Customer, error)
	Create(customer models.Customer) (string, error)
	SetAssignment(idCostumer string, newAssigment models.Assignment) error
	DeleteAll(experimentId string) (bool, error)
	DeleteAllOfOwner(owner models.User) (bool, error)
}
