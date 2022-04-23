package services

import (
	"lABoratory/lABoratoryAPI/internal/models"
	"testing"
)

func TestValidateExperiment(t *testing.T) {
	existingExperiments := []models.Experiment{
		{Id: "", Name: "Name 1", Description: "", Assignments: []models.Assignment{}, Owner: models.User{}},
		{Id: "", Name: "Name 2", Description: "", Assignments: []models.Assignment{}, Owner: models.User{}},
		{Id: "", Name: "Name 3", Description: "", Assignments: []models.Assignment{}, Owner: models.User{}},
	}
	validExperiment := models.Experiment{
		Id:          "",
		Name:        "Name",
		Description: "Description",
		Assignments: []models.Assignment{
			{AssignmentName: "a1", AssignmentValue: 25.0, AssignmentDescription: ""},
			{AssignmentName: "a2", AssignmentValue: 25.0, AssignmentDescription: ""},
			{AssignmentName: "a3", AssignmentValue: 25.0, AssignmentDescription: ""},
			{AssignmentName: "a4", AssignmentValue: 25.0, AssignmentDescription: ""},
		},
		Owner: models.User{
			Id:             "",
			Username:       "username 1",
			HashedPassword: "pw",
		},
	}
	invalidExperiment1 := models.Experiment{
		Id:          "",
		Name:        "Name",
		Description: "Description",
		Assignments: []models.Assignment{
			{AssignmentName: "Assignment 1", AssignmentValue: 25.0, AssignmentDescription: ""},
			{AssignmentName: "A2", AssignmentValue: 25.0, AssignmentDescription: ""},
			{AssignmentName: "a3", AssignmentValue: 25.0, AssignmentDescription: ""},
			{AssignmentName: "Assignment 4", AssignmentValue: 25.0, AssignmentDescription: ""},
		},
		Owner: models.User{
			Id:             "",
			Username:       "username 1",
			HashedPassword: "pw",
		},
	}
	invalidExperiment2 := models.Experiment{
		Id:          "",
		Name:        "Name",
		Description: "Description",
		Assignments: []models.Assignment{
			{AssignmentName: "a1", AssignmentValue: 25.0, AssignmentDescription: ""},
			{AssignmentName: "a2", AssignmentValue: 25.0, AssignmentDescription: ""},
			{AssignmentName: "a3", AssignmentValue: 0.0, AssignmentDescription: ""},
			{AssignmentName: "a4", AssignmentValue: 25.0, AssignmentDescription: ""},
		},
		Owner: models.User{
			Id:             "",
			Username:       "username 1",
			HashedPassword: "pw",
		},
	}
	invalidExperiment3 := models.Experiment{
		Id:          "",
		Name:        "Name 1",
		Description: "Description",
		Assignments: []models.Assignment{
			{AssignmentName: "a1", AssignmentValue: 25.0, AssignmentDescription: ""},
			{AssignmentName: "a2", AssignmentValue: 25.0, AssignmentDescription: ""},
			{AssignmentName: "a3", AssignmentValue: 25.0, AssignmentDescription: ""},
			{AssignmentName: "a4", AssignmentValue: 25.0, AssignmentDescription: ""},
		},
		Owner: models.User{
			Id:             "",
			Username:       "username 1",
			HashedPassword: "pw",
		},
	}
	es := ExperimentService{}
	valid := es.validateExperiment(validExperiment, existingExperiments)
	invalid1 := es.validateExperiment(invalidExperiment1, existingExperiments)
	invalid2 := es.validateExperiment(invalidExperiment2, existingExperiments)
	invalid3 := es.validateExperiment(invalidExperiment3, existingExperiments)
	if valid && !invalid1 && !invalid2 && !invalid3 {
		t.Log("ValidateExperiment PASSED")
	} else {
		t.Errorf("ValidateExperiment FAILED. Expected %t, %t, %t, %t, got%t, %t, %t, %t",
			true, false, false, false,
			valid, invalid1, invalid2, invalid3)
	}
}

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
