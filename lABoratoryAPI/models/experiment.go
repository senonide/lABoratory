package models

type Experiment struct {
	Id                string       `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string       `json:"name"`
	ActiveExperiments []assignment `json:"activeExperiments"`
}

type assignment struct {
	AssignmentName  string  `json:"assignmentName"`
	AssignmentValue float64 `json:"assignmentValue"`
}
