package services

import (
	"fmt"
	"lABoratory/lABoratoryAPI/models"
	"lABoratory/lABoratoryAPI/persistence"
	"lABoratory/lABoratoryAPI/utils"
)

type AuthService struct {
	repository        persistence.UserRepository
	securityProvider  utils.SecurityProviderI
	experimentService ExperimentServiceI
}

type AuthServiceI interface {
	GetAll() ([]models.User, error)
	GetOne(token string) (*models.User, error)
	Delete(token string) (bool, error)
	SignupUser(unknownUser models.User) (string, error)
	ValidateUser(unknownUser models.User) (string, error)
}

func NewAuthService(r persistence.UserRepository, sp utils.SecurityProviderI, es ExperimentServiceI) AuthServiceI {
	as := new(AuthService)
	as.repository = r
	as.securityProvider = sp
	as.experimentService = es
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
	jwtoken, err := as.securityProvider.GetToken(token)
	if err != nil {
		return nil, err
	}
	claims, err := as.securityProvider.GetTokenClaims(jwtoken)
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
	wereDeleted, err := as.experimentService.DeleteAll(user)
	if err != nil {
		return wasDeleted, fmt.Errorf("error deleting user experiments")
	}
	return wasDeleted && wereDeleted, nil
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
	return as.securityProvider.GenJWT(user)
}
