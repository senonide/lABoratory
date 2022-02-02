/*
	This class is in charge of interacting with the database, it applies the design pattern called fachade
*/
package database

import (
	"context"

	"lABoratory/lABoratoryAPI/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = GetCollection("experiments")
var ctx = context.Background()

func Create(experiment models.Experiment) error {

	var err error

	_, err = collection.InsertOne(ctx, experiment)

	if err != nil {
		return err
	}

	return nil
}

func Read() (models.AllExperiments, error) {

	var experiments models.AllExperiments

	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {

		var experiment models.Experiment
		err = cur.Decode(&experiment)

		if err != nil {
			return nil, err
		}

		experiments = append(experiments, experiment)
	}

	return experiments, nil
}

func ReadOne(experimentId string) (models.Experiment, error) {

	var experiment models.Experiment

	oid, _ := primitive.ObjectIDFromHex(experimentId)

	filter := bson.M{"_id": oid}

	cur := collection.FindOne(ctx, filter)

	err := cur.Decode(&experiment)

	if err != nil {
		return experiment, err
	}

	return experiment, nil
}

func Update(experiment models.Experiment, experimentId string) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(experimentId)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"name":              experiment.Name,
			"c":                 experiment.C,
			"activeExperiments": experiment.ActiveExperiments,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func Delete(experimentId string) error {

	var err error
	var oid primitive.ObjectID

	oid, err = primitive.ObjectIDFromHex(experimentId)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	_, err = collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
