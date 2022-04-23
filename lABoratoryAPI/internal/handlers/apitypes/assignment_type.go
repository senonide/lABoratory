package apitypes

type AssignmentType struct {
	AssignmentName        string  `json:"assignmentName"`
	AssignmentValue       float64 `json:"assignmentValue" binding:"required"`
	AssignmentDescription string  `json:"assignmentDescription"`
}
