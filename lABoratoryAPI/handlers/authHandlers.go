package handlers

import (
	"lABoratory/lABoratoryAPI/handlers/apitypes"
	"lABoratory/lABoratoryAPI/handlers/responses"
	"lABoratory/lABoratoryAPI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler() *AuthHandler {
	ah := new(AuthHandler)
	ah.service = services.NewAuthService()
	return ah
}

func (ah *AuthHandler) GetUsers(c *gin.Context) {
	users, err := ah.service.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, apitypes.GetUsersApiType(users))
}

func (ah *AuthHandler) GetUser(c *gin.Context) {
	var data map[string]string
	err := c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	user, err := ah.service.GetOne(data["jwt"])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	if user == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, apitypes.GetUserApiType(user))
}

func (ah *AuthHandler) DeleteUser(c *gin.Context) {
	var data map[string]string
	err := c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	wasDeleted, err := ah.service.Delete(data["jwt"])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	if !wasDeleted {
		c.AbortWithStatus(http.StatusNotModified)
		return
	}
	c.IndentedJSON(http.StatusOK, responses.DeleteResponse{WasDeleted: wasDeleted})
}

func (ah *AuthHandler) Authenticate(c *gin.Context) {
	var data map[string]string
	err := c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	unknownUser := apitypes.User{Username: data["username"], Password: data["password"]}
	token, err := ah.service.ValidateUser(unknownUser.GetUserModel())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, responses.ResponseWithError{Message: "error", Error: "incorrect password"})
		return
	}
	c.IndentedJSON(http.StatusOK, responses.ResponseWithToken{Message: "success", Token: token})
}

func (ah *AuthHandler) Signup(c *gin.Context) {
	var data map[string]string
	err := c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	unknownUser := apitypes.User{Username: data["username"], Password: data["password"]}
	token, err := ah.service.SignupUser(unknownUser.GetUserModel())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, responses.ResponseWithToken{Message: "user created successfully", Token: token})
}
