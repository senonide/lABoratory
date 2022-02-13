package types

type Experiment struct {
	Name              string       `json:"name"`
	ActiveExperiments []assignment `json:"activeExperiments"`
}

type assignment struct {
	AssignmentName  string  `json:"assignmentName"`
	AssignmentValue float64 `json:"assignmentValue"`
}
