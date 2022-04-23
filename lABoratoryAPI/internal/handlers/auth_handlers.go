package handlers

import (
	"lABoratory/lABoratoryAPI/internal/handlers/apitypes"
	"lABoratory/lABoratoryAPI/internal/handlers/responses"
	"lABoratory/lABoratoryAPI/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service services.AuthServiceI
}

func NewAuthHandler(as services.AuthServiceI) *AuthHandler {
	ah := new(AuthHandler)
	ah.service = as
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
	tokenFromHeader := c.Request.Header.Get("Authorization")
	user, err := ah.service.GetOne(tokenFromHeader)
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
	tokenFromHeader := c.Request.Header.Get("Authorization")
	wasDeleted, err := ah.service.Delete(tokenFromHeader)
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
	var data apitypes.User
	err := c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	token, err := ah.service.ValidateUser(data.GetUserModel())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, responses.ResponseWithError{Message: "error", Error: "incorrect password"})
		return
	}
	c.IndentedJSON(http.StatusCreated, responses.ResponseWithToken{Message: "success", Token: token})
}

func (ah *AuthHandler) Signup(c *gin.Context) {
	var data apitypes.User
	err := c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	token, err := ah.service.SignupUser(data.GetUserModel())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, responses.ResponseWithToken{Message: "success", Token: token})
}
