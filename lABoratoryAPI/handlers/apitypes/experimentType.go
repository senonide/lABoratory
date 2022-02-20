package apitypes

import "lABoratory/lABoratoryAPI/models"

type Experiment struct {
	Name        string              `json:"name"`
	Assignments []models.Assignment `json:"assignments"`
}

func GetExperimentApiType(exp models.Experiment) Experiment {
	return Experiment{Name: exp.Name, Assignments: exp.Assignments}
}

func (experiment Experiment) GetExperimentModel() models.Experiment {
	experimentModel := models.Experiment{Name: experiment.Name, Assignments: experiment.Assignments}
	return experimentModel
}
