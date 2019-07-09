package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GetToken(req *http.Request) (*string, error) {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("auth header empty")
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && (strings.ToLower(parts[0]) == "bearer")) {
		return nil, errors.New("Invalid authorization header")
	}

	token := parts[1]
	return &token, nil
}

func Validate(token, signedKey string) (*T1CGClaims, error) {

	parsed, err := jwt.ParseWithClaims(token, &T1CGClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method[%v]", token.Header["alg"])
		}
		return []byte(signedKey), nil
	})
	if err != nil {
		return nil, err
	}

	return parsed.Claims.(*T1CGClaims), nil
}

func IssueToken(userID string, role int64, issuer string, tokenExpiry time.Time, signedKey string, data interface{}) (string, error) {
	claims := T1CGClaims{
		userID,
		Role(role),
		data,
		jwt.StandardClaims{
			ExpiresAt: tokenExpiry.Unix(),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(signedKey))
	if err != nil {
		return "", errors.New("Failed to sign string")
	}

	return signedToken, nil
}
