package models

type Customer struct {
	Id                    string `bson:"_id,omitempty" json:"id,omitempty"`
	Key                   string `json:"key"`
	Override              bool   `json:"override"`
	ExperimentId          string `json:"experiment" binding:"required"`
	AssignmentName        string `json:"assignment" binding:"required"`
	AssignmentDescription string `json:"description,omitempty"`
}
