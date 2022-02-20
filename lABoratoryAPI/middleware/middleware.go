package middleware

import (
	"fmt"
	"lABoratory/lABoratoryAPI/handlers/apitypes"
	"lABoratory/lABoratoryAPI/handlers/responses"
	"lABoratory/lABoratoryAPI/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateJWT(c *gin.Context) {
	var token apitypes.Jwt
	err := c.BindJSON(&token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	jwtoken, err := utils.GetToken(token.Token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	if !utils.Validate(jwtoken) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responses.ResponseWithError{Message: "error", Error: fmt.Errorf("invalid token").Error()})
		return
	}
}
