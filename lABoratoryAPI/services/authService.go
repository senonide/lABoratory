package services

import (
	"fmt"
	"lABoratory/lABoratoryAPI/models"
	"lABoratory/lABoratoryAPI/persistence"
	"lABoratory/lABoratoryAPI/persistence/database"
	"lABoratory/lABoratoryAPI/utils"
)

type AuthService struct {
	repository persistence.UserRepository
}

func NewAuthService() *AuthService {
	as := new(AuthService)
	as.repository = database.NewDbUserRepository()
	return as
}

func (as *AuthService) GetAll() ([]models.User, error) {
	users, err := as.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (as *AuthService) GetOne(token string) (*models.User, error) {
	jwtoken, err := utils.GetToken(token)
	if err != nil {
		return nil, err
	}
	claims, err := utils.TokenClaims(jwtoken)
	if err != nil {
		return nil, err
	}
	usernameFromToken, ok := claims["sub"].(string)
	if !ok || usernameFromToken == "" {
		return nil, fmt.Errorf("token subject error")
	}
	user, err := as.repository.GetOne(usernameFromToken)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (as *AuthService) Delete(token string) (bool, error) {
	user, err := as.GetOne(token)
	if err != nil {
		return false, err
	}
	wasDeleted, err := as.repository.Delete(user.Id)
	if err != nil {
		return false, err
	}
	return wasDeleted, nil
}

func (as *AuthService) SignupUser(unknownUser models.User) (string, error) {
	_, err := as.repository.GetOne(unknownUser.Username)
	if err == nil {
		return "", fmt.Errorf("user already exists")
	}
	errCreating := as.repository.Create(unknownUser)
	if errCreating != nil {
		return "", errCreating
	}
	return as.ValidateUser(unknownUser)
}

func (as *AuthService) ValidateUser(unknownUser models.User) (string, error) {
	user, err := as.repository.GetOne(unknownUser.Username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", fmt.Errorf("user not found")
	}
	if user.HashedPassword != unknownUser.HashedPassword {
		return "", fmt.Errorf("incorrect password")
	}
	return utils.GenJWT(user)
}
