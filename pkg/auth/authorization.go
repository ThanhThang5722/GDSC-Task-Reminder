package auth

import (
	"crypto/rand"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	ID int
	jwt.StandardClaims
}

var (
	JWTKey []byte // todo: init
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
	claims.ExpiresAt = time.Now().Add(5 * time.Minute).Unix()

	// You can also optionally set other fields like IssuedAt or NotBefore
	claims.IssuedAt = time.Now().Unix()
	claims.NotBefore = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		return "", errors.New("fail to sign token")
	}

	return tokenString, nil
}
