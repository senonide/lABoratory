package database

import (
	"context"
	"fmt"
	"lABoratory/lABoratoryAPI/models"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const CostumersCollName string = "costumers"

type dbCustomerRepository struct {
	database *mongo.Database
}

func NewDbCustomerRepository() *dbCustomerRepository {
	repository := new(dbCustomerRepository)
	repository.database = GetDatabase()
	return repository
}

func (r *dbCustomerRepository) GetOne(customerId string) (*models.Customer, error) {
	ctx := context.Background()
	collection := r.database.Collection(CostumersCollName)
	var customer *models.Customer
	oid, _ := primitive.ObjectIDFromHex(customerId)
	filter := bson.M{"_id": oid}
	cur := collection.FindOne(ctx, filter)
	if cur.Err() != nil {
		return nil, nil
	}
	err := cur.Decode(&customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *dbCustomerRepository) Create(customer models.Customer) (string, error) {
	ctx := context.Background()
	collection := r.database.Collection(CostumersCollName)
	result, err := collection.InsertOne(ctx, customer)
	if err != nil {
		return "", err
	}
	id := fmt.Sprintf("%v", result)
	id = strings.Replace(id, "&{ObjectID(\"", "", 1)
	id = strings.Replace(id, "\")}", "", 1)
	return id, nil
}

func (r *dbCustomerRepository) SetAssignment(idCostumer string, newAssigment models.Assignment) error {
	ctx := context.Background()
	collection := r.database.Collection(CostumersCollName)
	oid, _ := primitive.ObjectIDFromHex(idCostumer)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"assignment":  newAssigment.AssignmentName,
			"description": newAssigment.AssignmentDescription,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *dbCustomerRepository) DeleteAll(experimentId string) (bool, error) {
	ctx := context.Background()
	collection := r.database.Collection(CostumersCollName)
	filter := bson.M{"experiment": experimentId}
	_, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *dbCustomerRepository) DeleteAllOfOwner(owner models.User) (bool, error) {
	return true, nil
}
