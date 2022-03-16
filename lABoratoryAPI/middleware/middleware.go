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
	tokenFromHeader := c.Request.Header.Get("Authorization")
	jwtoken, err := sp.GetToken(tokenFromHeader)
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
