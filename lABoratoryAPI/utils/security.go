package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"lABoratory/lABoratoryAPI/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type SecurityProvider struct{}

type SecurityProviderI interface {
	GetPasswordHash(password string) string
	GenJWT(subject string, expires bool) (string, error)
	GetToken(tokenString string) (*jwt.Token, error)
	GetTokenClaims(token *jwt.Token) (jwt.MapClaims, error)
	ValidateToken(token *jwt.Token) bool
}

func NewSecurityProvider() *SecurityProvider {
	sp := new(SecurityProvider)
	return sp
}

func (sp SecurityProvider) GetPasswordHash(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

func (sp SecurityProvider) GenJWT(subject string, expires bool) (string, error) {
	hmacSecret := []byte(config.ConfigParams.JwtSecret)
	var claims *jwt.RegisteredClaims
	if expires {
		var exp *jwt.NumericDate = new(jwt.NumericDate)
		exp.Time = time.Now().Add(time.Hour * 24)
		claims = &jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: exp,
		}
	} else {
		claims = &jwt.RegisteredClaims{
			Subject: subject,
		}
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(hmacSecret)
	return tokenString, err
}

func (sp SecurityProvider) GetToken(tokenString string) (*jwt.Token, error) {
	hmacSecret := []byte(config.ConfigParams.JwtSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})
	return token, err
}

func (sp SecurityProvider) ValidateToken(token *jwt.Token) bool {
	return token.Valid
}

func (sp SecurityProvider) GetTokenClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token format")
	}
}
