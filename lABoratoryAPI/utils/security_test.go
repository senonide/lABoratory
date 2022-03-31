package utils

import (
	"lABoratory/lABoratoryAPI/config"
	"lABoratory/lABoratoryAPI/models"
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

func TestGenJWT(t *testing.T) {
	securityProvider := NewSecurityProvider()
	testToken, err := securityProvider.GenJWT(&models.User{Id: "0", Username: "test", HashedPassword: "test"})
	if err != nil {
		t.Errorf("Security Test (GenJWT) FAILED. Error generating test token: %s", err)
	}
	hmacSecret := []byte(config.ConfigParams.JwtSecret)
	token, err := jwt.Parse(testToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			t.Errorf("Security Test (GenJWT) FAILED. Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})
	if err != nil {
		t.Errorf("Security Test (GenJWT) FAILED. Error parsing test token: %s", err)
	}
	if !token.Valid || err != nil {
		t.Errorf("Security Test (GenJWT) FAILED. %s", err)
	} else {
		t.Logf("Security Test (GenJWT) PASSED")
	}
}

func TestGetToken(t *testing.T) {
	securityProvider := NewSecurityProvider()
	hmacSecret := []byte(config.ConfigParams.JwtSecret)
	claims := &jwt.RegisteredClaims{
		Subject: "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newToken, err := token.SignedString(hmacSecret)
	if err != nil {
		t.Errorf("Security Test (GetToken) FAILED. Error generating test token: %s", err)
	}
	testToken, err := securityProvider.GetToken(newToken)
	if err != nil || !testToken.Valid {
		t.Errorf("Security Test (GetToken) FAILED. Error: %s", err)
	} else {
		t.Logf("Security Test (GetToken) PASSED")
	}
}

func TestGetTokenClaims(t *testing.T) {
	securityProvider := NewSecurityProvider()
	hmacSecret := []byte(config.ConfigParams.JwtSecret)
	claims := &jwt.RegisteredClaims{
		Subject: "test",
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenSigned, err := newToken.SignedString(hmacSecret)
	if err != nil {
		t.Errorf("Security Test (GetTokenClaims) FAILED. Error generating test token: %s", err)
	}
	testToken, err := jwt.Parse(newTokenSigned, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			t.Errorf("Security Test (GetTokenClaims) FAILED. Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})
	if err != nil {
		t.Errorf("Security Test (GetTokenClaims) FAILED. Error parsing test token: %s", err)
	}
	testClaims, err := securityProvider.GetTokenClaims(testToken)
	if err != nil {
		t.Errorf("Security Test (GetTokenClaims) FAILED. Error getting claims: %s", err)
	}
	if jwt.Claims(testClaims).Valid() != jwt.Claims(claims).Valid() {
		t.Errorf("Security Test (GetTokenClaims) FAILED. Claims Does not match")
	} else {
		t.Logf("Security Test (GetTokenClaims) PASSED")
	}
}
