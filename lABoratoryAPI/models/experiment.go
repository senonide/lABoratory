package models

import (
	"fmt"
	"math/rand"
)

type Experiment struct {
	Id          string       `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string       `json:"name" binding:"required"`
	Description string       `json:"description,omitempty"`
	Assignments []Assignment `json:"assignments" binding:"required"`
	Owner       User         `json:"owner,omitempty"`
}

func (exp Experiment) GetAssignmentByName(assignmentName string) (*Assignment, error) {
	for _, assignment := range exp.Assignments {
		if assignment.AssignmentName == assignmentName {
			return &assignment, nil
		}
	}
	return nil, fmt.Errorf("the assignment does not exists")
}

func (exp Experiment) GetRandomAssignment() Assignment {
	index := rand.Intn(len(exp.Assignments) - 1)
	return exp.Assignments[index]
}
