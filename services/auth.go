package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
)

// Auth service relating to authorization
type Auth struct {
	env    lib.Env
	logger lib.Logger
}

// NewAuth creates a new auth service
func NewAuth(env lib.Env, logger lib.Logger) Auth {
	return Auth{
		env:    env,
		logger: logger,
	}
}

// Authorize authorizes the generated token
func (s Auth) Authorize(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.env.JWTSecret), nil
	})
	if token.Valid {
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New("token malformed")
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New("token expired")
		}
	}
	return false, errors.New("couldn't handle token")
}

// CreateToken creates jwt auth token
func (s Auth) CreateToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.ID,
		"username":   user.Username,
		"created_at": time.Now().UnixMilli(),
	})

	tokenString, err := token.SignedString([]byte(s.env.JWTSecret))

	if err != nil {
		s.logger.Error("JWT validation failed: ", err)
	}

	return tokenString
}