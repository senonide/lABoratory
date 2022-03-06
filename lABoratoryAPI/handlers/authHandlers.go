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

func NewAuthHandler(as *services.AuthService) *AuthHandler {
	ah := new(AuthHandler)
	ah.service = as
	return ah
}

func (ah *AuthHandler) GetUsers(c *gin.Context) {
	users, err := ah.service.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, apitypes.GetUsersApiType(users))
}

func (ah *AuthHandler) GetUser(c *gin.Context) {
	tokenFromCookie, err := c.Cookie("jwt")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	user, err := ah.service.GetOne(tokenFromCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	if user == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, apitypes.GetUserApiType(user))
}

func (ah *AuthHandler) DeleteUser(c *gin.Context) {
	tokenFromCookie, err := c.Cookie("jwt")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	wasDeleted, err := ah.service.Delete(tokenFromCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
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
	c.SetCookie("jwt", token, 1, "/", "localhost", true, true)
	c.IndentedJSON(http.StatusCreated, responses.Response{Message: "success"})
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
	c.SetCookie("jwt", token, 1, "/", "localhost", true, true)
	c.IndentedJSON(http.StatusCreated, responses.Response{Message: "user created successfully"})
}
