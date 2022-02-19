package persistence

import "lABoratory/lABoratoryAPI/models"

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetOne(username string) (*models.User, error)
	Create(user models.User) error
	Update(user models.User) error
	Delete(userId string) (bool, error)
}
