package models

type Experiment struct {
	Id                string       `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string       `json:"name"`
	ActiveExperiments []Assignment `json:"activeExperiments"`
}
