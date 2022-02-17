package database

import (
	"context"
	"lABoratory/lABoratoryAPI/config"
	"lABoratory/lABoratoryAPI/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type dbExperimentRepository struct {
	database *mongo.Database
}

func NewDbExperimentRepository() *dbExperimentRepository {
	repository := new(dbExperimentRepository)
	repository.database = GetDatabase()
	return repository
}

func (r *dbExperimentRepository) GetAll() ([]models.Experiment, error) {
	ctx := context.Background()
	collection := r.database.Collection(config.ConfigParams.ExperimentCollName)
	var experiments []models.Experiment
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

func (r *dbExperimentRepository) GetOne(experimentId string) (*models.Experiment, error) {
	ctx := context.Background()
	collection := r.database.Collection(config.ConfigParams.ExperimentCollName)
	var experiment *models.Experiment
	oid, _ := primitive.ObjectIDFromHex(experimentId)
	filter := bson.M{"_id": oid}
	cur := collection.FindOne(ctx, filter)
	err := cur.Decode(&experiment)
	if err != nil {
		return nil, err
	}
	return experiment, nil
}

func (r *dbExperimentRepository) Create(experiment models.Experiment) error {
	ctx := context.Background()
	collection := r.database.Collection(config.ConfigParams.ExperimentCollName)
	_, err := collection.InsertOne(ctx, experiment)
	if err != nil {
		return err
	}
	return nil
}

func (r *dbExperimentRepository) Update(experiment models.Experiment) error {
	ctx := context.Background()
	collection := r.database.Collection(config.ConfigParams.ExperimentCollName)
	oid, _ := primitive.ObjectIDFromHex(experiment.Id)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"name":              experiment.Name,
			"activeExperiments": experiment.ActiveExperiments,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *dbExperimentRepository) Delete(experimentId string) (bool, error) {
	ctx := context.Background()
	collection := r.database.Collection(config.ConfigParams.ExperimentCollName)
	oid, errDecoding := primitive.ObjectIDFromHex(experimentId)
	if errDecoding != nil {
		return false, errDecoding
	}
	filter := bson.M{"_id": oid}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}
