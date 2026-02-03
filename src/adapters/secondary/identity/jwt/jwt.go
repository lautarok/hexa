package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/domain/models"
	"github.com/lautarok/hexa/src/domain/ports"
)

type JWTAdapter struct {
	secret string
}

type JWTAdapterDeps struct {
	Secret string
}

func NewJWTAdapter(deps *JWTAdapterDeps) ports.IdentityPort {
	return &JWTAdapter{
		secret: deps.Secret,
	}
}

func (adapter *JWTAdapter) NewToken(identity *models.Identity) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 72).Unix()

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subUserId": identity.SubUserID,
		"userId":    identity.UserID,
		"exp":       exp,
	})

	token, err := jwtToken.SignedString([]byte(adapter.secret))
	if err != nil {
		return "", 0, err
	}

	return token, exp, nil
}

func (adapter *JWTAdapter) ParseToken(token string) (*models.Identity, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(adapter.secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok {
		userId, err := uuid.Parse(claims["userId"].(string))
		if err != nil {
			return nil, err
		}

		subUserId, err := uuid.Parse(claims["subUserId"].(string))
		if err != nil {
			return nil, err
		}

		return &models.Identity{
			SubUserID: subUserId,
			UserID:    userId,
		}, nil
	}

	return nil, errors.New("Invalid claims")
}
