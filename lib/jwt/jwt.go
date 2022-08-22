package jwt

import (
	"errors"
	"time"

	"github.com/Augenblick-tech/bilibot/lib/conf"
	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID         uint   `json:"user_id"`
	Username       string `json:"username"`
	IsRefreshToken bool   `json:"is_refresh_token"`
	jwt.RegisteredClaims
}

const (
	TokenExpireDuration   = time.Hour * 2
	ReTokenExpireDuration = time.Hour * 24 * 3
)

var Secret = []byte(conf.C.JWT.Secret)

func GenToken(userID uint, username string) (string, error) {
	return generate(userID, username, false, TokenExpireDuration)
}

func GenReToken(userID uint, username string) (string, error) {
	return generate(userID, username, true, ReTokenExpireDuration)
}

func generate(userID uint, username string, IsRefreshToken bool, ExpireTime time.Duration) (string, error) {
	claims := Claims{
		userID,
		username,
		IsRefreshToken,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ExpireTime)),
			Issuer:    "bilibot",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(Secret)
	if err != nil {
		return "", err
	}
	return "Bearer " + tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	if tokenString[0:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, e.ErrTokenExpired
		}
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
