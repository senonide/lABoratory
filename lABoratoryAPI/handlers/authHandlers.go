package handlers

import (
	"lABoratory/lABoratoryAPI/models"
	"lABoratory/lABoratoryAPI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.AuthService
}

type IAuthHandler interface {
	Authenticate(*gin.Context)
	Singin(*gin.Context)
}

func NewAuthHandler() *AuthHandler {
	ah := new(AuthHandler)
	ah.service = services.NewAuthService()
	return ah
}

func (ah *AuthHandler) Authenticate(c *gin.Context) {
	var credentials models.Credentials
	err := c.BindJSON(&credentials)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	user, errUserNotFound := ah.service.ValidateUser(credentials)
	if errUserNotFound != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, errUserNotFound.Error())
		return
	}
	token, errGenerating := ah.service.GenJWT(user)
	if errGenerating != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errGenerating.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, token)
}

func (ah *AuthHandler) Singup(c *gin.Context) {
	var credentials models.Credentials
	err := c.BindJSON(&credentials)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	user, errCreating := ah.service.SingupUser(credentials)
	if errCreating != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errCreating.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, user)
}
