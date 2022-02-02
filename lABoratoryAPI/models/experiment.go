package models

// Experiment Model
type Experiment struct {
	Id                string `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string `json:"name"`
	C                 int    `json:"c"`
	ActiveExperiments []int  `json:"activeExperiments"`
}

type AllExperiments []Experiment
