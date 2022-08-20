package jwt

import (
	"time"

	"github.com/Augenblick-tech/bilibot/lib/conf"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

const (
	TokenExpireDuration = time.Hour * 2
	ReTokenExpireDuration = time.Hour * 24 * 3
)

var Secret = []byte(conf.C.JWT.Secret)

func GenToken(userID uint, username string) (string, error) {
	return generate(userID, username, TokenExpireDuration)
}

func GenReToken(userID uint, username string) (string, error) {
	return generate(userID, username, ReTokenExpireDuration)
}

func generate(userID uint, username string, ExpireTime time.Duration) (string, error) {
	claims := Claims{
		userID,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ExpireTime)),
			Issuer:    "bilibot",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret)
}

func ParseToken(tokenString string) (*Claims, error) {
	if tokenString[0:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
