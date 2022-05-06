package apitypes

import (
	"lABoratory/lABoratoryAPI/internal/models"
	"lABoratory/lABoratoryAPI/internal/utils"
	"strconv"
)

type Experiment struct {
	Id            string           `json:"id"`
	Name          string           `json:"name" binding:"required"`
	Description   string           `json:"description"`
	ExperimentKey string           `json:"experimentKey"`
	Assignments   []AssignmentType `json:"assignments" binding:"required"`
}

func GetExperimentApiType(exp models.Experiment) Experiment {
	assignmentsType := []AssignmentType{}
	for _, assignment := range exp.Assignments {
		assignmentsType = append(assignmentsType, AssignmentType{
			AssignmentName:        assignment.AssignmentName,
			AssignmentValue:       assignment.AssignmentValue,
			AssignmentDescription: assignment.AssignmentDescription,
		})
	}
	return Experiment{
		Id:            exp.Id,
		Name:          exp.Name,
		Description:   exp.Description,
		ExperimentKey: getExperimentKey(exp),
		Assignments:   assignmentsType,
	}
}

func GetExperimentsApiType(experiments []models.Experiment) []Experiment {
	experimentsType := []Experiment{}
	for _, experiment := range experiments {
		experimentsType = append(experimentsType, GetExperimentApiType(experiment))
	}
	return experimentsType
}

func (experiment Experiment) GetExperimentModel() models.Experiment {
	assignmentsModel := []models.Assignment{{
		AssignmentName:        "c",
		AssignmentValue:       experiment.Assignments[0].AssignmentValue,
		AssignmentDescription: experiment.Assignments[0].AssignmentDescription,
	}}
	for i := 1; i < len(experiment.Assignments); i++ {
		assignmentsModel = append(assignmentsModel, models.Assignment{
			AssignmentName:        "a" + strconv.Itoa(i),
			AssignmentValue:       experiment.Assignments[i].AssignmentValue,
			AssignmentDescription: experiment.Assignments[i].AssignmentDescription,
		})
	}
	experimentModel := models.Experiment{Name: experiment.Name, Description: experiment.Description, Assignments: assignmentsModel}
	if experiment.Id != "" {
		experimentModel.Id = experiment.Id
	}
	return experimentModel
}

func getExperimentKey(experiment models.Experiment) string {
	sp := new(utils.SecurityProvider)
	key, err := sp.GenJWT(experiment.Id, false)
	if err != nil {
		return ""
	}
	return key
}
