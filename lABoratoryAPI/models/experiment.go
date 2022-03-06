package models

type Experiment struct {
	Id          string       `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string       `json:"name" binding:"required"`
	Assignments []Assignment `json:"assignments" binding:"required"`
	Owner       User         `json:"owner,omitempty"`
}
