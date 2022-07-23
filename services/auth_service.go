package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go-clean-blog/lib"
	"go-clean-blog/models"

	"github.com/dgrijalva/jwt-go"
)

type JWTToken struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
}

type claims struct {
	jwt.StandardClaims
	JWTToken
}

// AuthService service relating to authorization
type AuthService struct {
	env    lib.Env
	logger lib.Logger
}

// NewAuthService creates a new auth service
func NewAuthService(env lib.Env, logger lib.Logger) AuthService {
	return AuthService{
		env:    env,
		logger: logger,
	}
}

// Authorize authorizes the generated token
func (s AuthService) Authorize(tokenString string) (*claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.env.JWTSecret), nil
	})
	if token.Valid {
		claims := token.Claims.(*claims)
		return claims, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.New("token malformed")
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, errors.New("token expired")
		}
	}
	return nil, errors.New("couldn't handle token")
}

// CreateToken creates jwt auth token
func (s AuthService) CreateToken(user models.User) string {
	t := JWTToken{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: fmt.Sprint(time.Now().UnixMilli()),
	}
	data, _ := json.Marshal(t)
	var claims jwt.MapClaims
	json.Unmarshal(data, &claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.env.JWTSecret))

	if err != nil {
		s.logger.Error("JWT validation failed: ", err)
	}

	return tokenString
}
