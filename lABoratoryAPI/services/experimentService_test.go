package services

import (
	"lABoratory/lABoratoryAPI/models"
	"testing"
)

func TestContainsExperimentName(t *testing.T) {
	experiments := []models.Experiment{
		{Id: "", Name: "Name 1", Description: "", Assignments: []models.Assignment{}, Owner: models.User{}},
		{Id: "", Name: "Name 2", Description: "", Assignments: []models.Assignment{}, Owner: models.User{}},
		{Id: "", Name: "Name 3", Description: "", Assignments: []models.Assignment{}, Owner: models.User{}},
	}
	valid := !containsExperimentName(experiments, "Name 4")
	invalid := !containsExperimentName(experiments, "Name 1")
	if valid && !invalid {
		t.Log("ContainsExperimentName PASSED")
	} else {
		t.Errorf("ContainsExperimentName FAILED. Expected %t and %t, got %t and %t", true, false, valid, invalid)
	}
}

func TestIsDuplicated(t *testing.T) {
	validAssignments := []models.Assignment{
		{AssignmentName: "Assignment 1", AssignmentValue: 25.0, AssignmentDescription: ""},
		{AssignmentName: "Assignment 2", AssignmentValue: 25.0, AssignmentDescription: ""},
		{AssignmentName: "Assignment 3", AssignmentValue: 25.0, AssignmentDescription: ""},
		{AssignmentName: "Assignment 4", AssignmentValue: 25.0, AssignmentDescription: ""},
	}
	invalidAssignments := []models.Assignment{
		{AssignmentName: "Assignment 1", AssignmentValue: 25.0, AssignmentDescription: ""},
		{AssignmentName: "Assignment 2", AssignmentValue: 25.0, AssignmentDescription: ""},
		{AssignmentName: "Assignment 1", AssignmentValue: 25.0, AssignmentDescription: ""},
	}
	valid := !isDuplicated(validAssignments)
	invalid := !isDuplicated(invalidAssignments)
	if valid && !invalid {
		t.Log("IsDuplicated PASSED")
	} else {
		t.Errorf("IsDuplicated FAILED. Expected %t and %t, got %t and %t", true, false, valid, invalid)
	}
}

func TestValidateOwnership(t *testing.T) {
	validOwner := models.User{Id: "", Username: "username 1", HashedPassword: "pw"}
	invalidOwner := models.User{Id: "", Username: "username 2", HashedPassword: "pw"}
	experiment := models.Experiment{
		Id:          "",
		Name:        "Name",
		Description: "Description",
		Assignments: []models.Assignment{},
		Owner: models.User{
			Id:             "",
			Username:       "username 1",
			HashedPassword: "pw",
		},
	}
	valid := validateOwnership(&experiment, &validOwner)
	invalid := validateOwnership(&experiment, &invalidOwner)
	if valid && !invalid {
		t.Log("ValidateOwnership PASSED")
	} else {
		t.Errorf("ValidateOwnership FAILED. Expected %t and %t, got %t and %t", true, false, valid, invalid)
	}
}
