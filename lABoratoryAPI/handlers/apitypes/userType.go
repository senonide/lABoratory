package apitypes

import (
	"lABoratory/lABoratoryAPI/models"
	"lABoratory/lABoratoryAPI/utils"
)

type User struct {
	Username         string `json:"username"`
	UnhashedPassword string `json:"password"`
}

func (user User) GetUserModel() models.User {
	passwordHash := utils.GetPasswordHash(user.UnhashedPassword)
	userModel := models.User{Username: user.Username, HashedPassword: passwordHash}
	return userModel
}
