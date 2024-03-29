package apitypes

import (
	"lABoratory/lABoratoryAPI/internal/models"
	"lABoratory/lABoratoryAPI/internal/utils"
)

type User struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetUserApiType(userModel *models.User) User {
	return User{Id: userModel.Id, Username: userModel.Username, Password: userModel.HashedPassword}
}

func GetUsersApiType(users []models.User) []User {
	UsersType := []User{}
	for _, user := range users {
		UsersType = append(UsersType, GetUserApiType(&user))
	}
	return UsersType
}

func (user User) GetUserModel() models.User {
	sp := new(utils.SecurityProvider)
	passwordHash := sp.GetPasswordHash(user.Password)
	userModel := models.User{Username: user.Username, HashedPassword: passwordHash}
	return userModel
}
