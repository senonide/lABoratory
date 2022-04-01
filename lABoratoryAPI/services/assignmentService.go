package services

import "lABoratory/lABoratoryAPI/models"

type AssignmentService struct{}

type AssignmentServiceI interface {
	GetAssignment(token string) (*models.Customer, error)
}

func NewAssignmentService() AssignmentServiceI {
	as := new(AssignmentService)
	return as
}

func (as AssignmentService) GetAssignment(token string) (*models.Customer, error) {
	//TODO: Diferenciate between creating a new assignment or returning the existing one
	return nil, nil
}
