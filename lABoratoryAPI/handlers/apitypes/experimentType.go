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

func GetExperimentsApiType(experiments []models.Experiment) []Experiment {
	experimentsType := []Experiment{}
	for _, experiment := range experiments {
		experimentsType = append(experimentsType, GetExperimentApiType(experiment))
	}
	return experimentsType
}

func (experiment Experiment) GetExperimentModel() models.Experiment {
	experimentModel := models.Experiment{Name: experiment.Name, Assignments: experiment.Assignments}
	return experimentModel
}
