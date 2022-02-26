package middleware

import (
	"fmt"
	"lABoratory/lABoratoryAPI/handlers/responses"
	"lABoratory/lABoratoryAPI/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateJWT(c *gin.Context) {
	sp := new(utils.SecurityProvider)
	tokenFromCookie, err := c.Cookie("jwt")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: "jwt not founded"})
		return
	}
	jwtoken, err := sp.GetToken(tokenFromCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: "unauthorized"})
		return
	}
	if !sp.ValidateToken(jwtoken) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responses.ResponseWithError{Message: "error", Error: fmt.Errorf("invalid token").Error()})
		return
	}
	c.Next()
}
