/*
	This class is in charge of implementing the CRUD methods
*/
package services

import (
	"lABoratory/lABoratoryAPI/database"
	"lABoratory/lABoratoryAPI/models"
)

func Create(experiment models.Experiment) error {

	err := database.Create(experiment)

	if err != nil {
		return err
	}

	return nil
}

func Read() (models.AllExperiments, error) {

	experiments, err := database.Read()

	if err != nil {
		return nil, err
	}

	return experiments, nil
}

func ReadOne(experimentId string) (models.Experiment, error) {

	experiment, err := database.ReadOne(experimentId)

	if err != nil {
		return experiment, err
	}

	return experiment, nil
}

func Update(experiment models.Experiment, experimentId string) error {

	err := database.Update(experiment, experimentId)

	if err != nil {
		return err
	}

	return nil
}

func Delete(experimentId string) error {

	err := database.Delete(experimentId)

	if err != nil {
		return err
	}

	return nil
}
