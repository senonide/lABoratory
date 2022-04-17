package services

import (
	"lABoratory/lABoratoryAPI/models"
	"testing"
)

func TestGetNewBalancedAssignment(t *testing.T) {
	as := AssignmentService{}
	a1 := models.Assignment{AssignmentName: "Assignment 1", AssignmentValue: 25.0, AssignmentDescription: "Assignment description"}
	a2 := models.Assignment{AssignmentName: "Assignment 2", AssignmentValue: 25.0, AssignmentDescription: "Assignment description"}
	a3 := models.Assignment{AssignmentName: "Assignment 3", AssignmentValue: 25.0, AssignmentDescription: "Assignment description"}
	a4 := models.Assignment{AssignmentName: "Assignment 4", AssignmentValue: 25.0, AssignmentDescription: "Assignment description"}
	experiment := models.Experiment{
		Id:          "id",
		Name:        "Name",
		Description: "Description",
		Assignments: []models.Assignment{a1, a2, a3, a4},
		Owner: models.User{
			Id:             "",
			Username:       "username",
			HashedPassword: "pw",
		},
	}
	previousAssignments := []models.Customer{
		{Id: "", ExperimentId: "id", AssignmentName: a1.AssignmentName, AssignmentDescription: a1.AssignmentDescription},
		{Id: "", ExperimentId: "id", AssignmentName: a2.AssignmentName, AssignmentDescription: a2.AssignmentDescription},
		{Id: "", ExperimentId: "id", AssignmentName: a4.AssignmentName, AssignmentDescription: a4.AssignmentDescription},
	}
	newAssignment, err := as.getNewBalancedAssignment(&experiment, &previousAssignments)
	if err != nil {
		t.Errorf("GetNewBalancedAssignment FAILED. Got an error %s", err)
	}
	if newAssignment.AssignmentName != a3.AssignmentName {
		t.Errorf("GetNewBalancedAssignment FAILED. Expected %s, got %s", a3.AssignmentName, newAssignment.AssignmentName)
	} else {
		t.Log("GetNewBalancedAssignment PASSED")
	}

}
