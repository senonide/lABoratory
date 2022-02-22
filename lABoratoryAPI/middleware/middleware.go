package middleware

import (
	"fmt"
	"lABoratory/lABoratoryAPI/handlers/responses"
	"lABoratory/lABoratoryAPI/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateJWT(c *gin.Context) {
	var data map[string]interface{}
	err := c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responses.ResponseWithError{Message: "error", Error: "unauthorized"})
		return
	}
	jwtoken, err := utils.GetToken(data["jwt"].(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: "unauthorized"})
		return
	}
	if !utils.Validate(jwtoken) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responses.ResponseWithError{Message: "error", Error: fmt.Errorf("invalid token").Error()})
		return
	}
	c.Next()
}
