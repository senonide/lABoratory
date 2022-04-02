package apitypes

import (
	"lABoratory/lABoratoryAPI/models"
	"lABoratory/lABoratoryAPI/utils"
)

type Experiment struct {
	Id            string              `json:"id"`
	Name          string              `json:"name" binding:"required"`
	Description   string              `json:"description"`
	ExperimentKey string              `json:"experimentKey"`
	Assignments   []models.Assignment `json:"assignments" binding:"required"`
}

func GetExperimentApiType(exp models.Experiment) Experiment {
	return Experiment{Id: exp.Id, Name: exp.Name, Description: exp.Description, ExperimentKey: getExperimentKey(exp), Assignments: exp.Assignments}
}

func GetExperimentsApiType(experiments []models.Experiment) []Experiment {
	experimentsType := []Experiment{}
	for _, experiment := range experiments {
		experimentsType = append(experimentsType, GetExperimentApiType(experiment))
	}
	return experimentsType
}

func (experiment Experiment) GetExperimentModel() models.Experiment {
	experimentModel := models.Experiment{Name: experiment.Name, Description: experiment.Description, Assignments: experiment.Assignments}
	return experimentModel
}

func getExperimentKey(experiment models.Experiment) string {
	sp := new(utils.SecurityProvider)
	key, err := sp.GenJWT(experiment.Id, false)
	if err != nil {
		return ""
	}
	return "experimentKey_" + key
}
