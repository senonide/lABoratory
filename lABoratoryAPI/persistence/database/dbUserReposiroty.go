package database

import (
	"lABoratory/lABoratoryAPI/models"

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
	return nil, nil
}

func (r *dbUserRepository) GetOne(username string) (*models.User, error) {
	return nil, nil
}

func (r *dbUserRepository) Create(user models.User) error {
	return nil
}

func (r *dbUserRepository) Update(user models.User) error {
	return nil
}

func (r *dbUserRepository) Delete(userId string) (bool, error) {
	return false, nil
}
