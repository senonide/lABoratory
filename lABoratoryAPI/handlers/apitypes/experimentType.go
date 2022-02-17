package apitypes

import "lABoratory/lABoratoryAPI/models"

type Experiment struct {
	Id                string              `json:"id"`
	Name              string              `json:"name"`
	ActiveExperiments []models.Assignment `json:"activeExperiments"`
}

func GetExperimentApiType(exp models.Experiment) Experiment {
	return Experiment{Id: exp.Id, Name: exp.Name, ActiveExperiments: exp.ActiveExperiments}
}
