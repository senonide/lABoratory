package services

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"lABoratory/lABoratoryAPI/config"
	"lABoratory/lABoratoryAPI/models"
	"lABoratory/lABoratoryAPI/persistence"
	"lABoratory/lABoratoryAPI/persistence/database"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	repository persistence.UserRepository
}

func NewAuthService() *AuthService {
	as := new(AuthService)
	as.repository = database.NewDbUserRepository()
	return as
}

func (as *AuthService) GenJWT(user *models.User) (string, error) {
	hmacSecret := []byte(config.ConfigParams.JwtSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 6).Unix(),
	})
	tokenString, err := token.SignedString(hmacSecret)
	return tokenString, err
}

func (as *AuthService) ValidateJWT(tokenString string) (bool, jwt.Claims, error) {
	hmacSecret := []byte(config.ConfigParams.JwtSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return true, claims, nil
	} else {
		return false, nil, err
	}
}

func (as *AuthService) SingupUser(credentials models.Credentials) (*models.User, error) {
	hashedPassword := hashPassword(credentials.Password)
	credentials.Password = hashedPassword
	// TODO: Create with the persistance repository a new user
	return nil, nil
}

func (as *AuthService) ValidateUser(credentials models.Credentials) (*models.User, error) {
	hashedPassword := hashPassword(credentials.Password)
	credentials.Password = hashedPassword
	// TODO: Get with the persistance repository the user that match with the given credentials
	return nil, nil
}

func hashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}
