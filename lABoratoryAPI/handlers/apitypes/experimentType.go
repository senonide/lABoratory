package apitypes

import "lABoratory/lABoratoryAPI/models"

type Experiment struct {
	Id          string              `json:"id"`
	Name        string              `json:"name"`
	Assignments []models.Assignment `json:"assignments"`
}

func GetExperimentApiType(exp models.Experiment) Experiment {
	return Experiment{Id: exp.Id, Name: exp.Name, Assignments: exp.Assignments}
}
