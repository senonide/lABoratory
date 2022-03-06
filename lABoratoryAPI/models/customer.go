package models

type Customer struct {
	Id             string `bson:"_id,omitempty" json:"id,omitempty"`
	ExperimentId   string `json:"experiment" binding:"required"`
	AssignmentName string `json:"assignment" binding:"required"`
}
