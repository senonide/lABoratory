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
	c.IndentedJSON(http.StatusOK, users)
}

func (ah *AuthHandler) GetUser(c *gin.Context) {
	var token apitypes.Jwt
	err := c.BindJSON(&token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	user, err := ah.service.GetOne(token.Token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	if user == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (ah *AuthHandler) DeleteUser(c *gin.Context) {
	var token apitypes.Jwt
	err := c.BindJSON(&token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	wasDeleted, err := ah.service.Delete(token.Token)
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
	var unknownUser apitypes.User
	err := c.BindJSON(&unknownUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	token, err := ah.service.ValidateUser(unknownUser.GetUserModel())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, responses.ResponseWithToken{Message: "success", Token: token})
}

func (ah *AuthHandler) Signup(c *gin.Context) {
	var unknownUser apitypes.User
	err := c.BindJSON(&unknownUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	token, err := ah.service.SignupUser(unknownUser.GetUserModel())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, responses.ResponseWithToken{Message: "user created successfully", Token: token})
}
