package models

type Assignment struct {
	AssignmentName        string  `json:"assignmentName" binding:"required"`
	AssignmentValue       float64 `json:"assignmentValue" binding:"required"`
	AssignmentDescription string  `json:"assignmentDescription"`
}
