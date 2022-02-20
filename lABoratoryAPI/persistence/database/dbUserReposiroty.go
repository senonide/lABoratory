package database

import (
	"context"
	"lABoratory/lABoratoryAPI/config"
	"lABoratory/lABoratoryAPI/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type dbUserRepository struct {
	database *mongo.Database
}

func NewDbUserRepository() *dbUserRepository {
	repository := new(dbUserRepository)
	repository.database = GetDatabase()
	return repository
}

func (r *dbUserRepository) GetAll() ([]models.User, error) {
	ctx := context.Background()
	collection := r.database.Collection(config.ConfigParams.UsersCollName)
	users := []models.User{}
	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var user models.User
		err = cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *dbUserRepository) GetOne(username string) (*models.User, error) {
	ctx := context.Background()
	collection := r.database.Collection(config.ConfigParams.UsersCollName)
	var user *models.User
	filter := bson.M{"username": username}
	cur := collection.FindOne(ctx, filter)
	if cur.Err() != nil {
		return nil, cur.Err()
	}
	err := cur.Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *dbUserRepository) Create(user models.User) error {
	ctx := context.Background()
	collection := r.database.Collection(config.ConfigParams.UsersCollName)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *dbUserRepository) Update(user models.User) error {
	ctx := context.Background()
	collection := r.database.Collection(config.ConfigParams.UsersCollName)
	oid, _ := primitive.ObjectIDFromHex(user.Id)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"username": user.Username,
			"password": user.HashedPassword,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *dbUserRepository) Delete(userId string) (bool, error) {
	ctx := context.Background()
	collection := r.database.Collection(config.ConfigParams.UsersCollName)
	oid, errDecoding := primitive.ObjectIDFromHex(userId)
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
