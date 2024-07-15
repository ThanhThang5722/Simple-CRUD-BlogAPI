package auth

import (
	"crypto/rand"
	"errors"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

var (
	JWTKey []byte
)

func GenerateJWTKey() {
	JWTKey = make([]byte, 32)
	rand.Read(JWTKey)
	log.Println(JWTKey)
}

func GetTokenString(ctx *gin.Context) string {
	var tokenString string
	authHeader := ctx.Request.Header.Get("Authorization")
	parts := strings.Split(authHeader, " ")
	if parts[0] == "Bearer" {
		tokenString = parts[1]
	}
	return tokenString
}

func ParseToken(tokenString string) (Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return Claims{}, errors.New("signature invalid")
		} else {
			return Claims{}, errors.New("can't parse token string")
		}
	}
	if !token.Valid {
		return Claims{}, errors.New("token invalid")
	}
	return *claims, nil
}

func GenerateTokenString(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		return "", errors.New("fail to sign token")
	}
	return tokenString, nil
}
