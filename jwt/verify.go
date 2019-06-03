package jwt

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GetToken
// params
//  req: a http request
// returns a token from the requests headers
func GetToken(req *http.Request) (*string, error) {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("auth header empty")
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && (strings.ToLower(parts[0]) == "bearer")) {
		return nil, fmt.Errorf("invalid auth header")
	}

	token := parts[1]
	return &token, nil
}

// Validate
// param
//  token: jwt token as a string
//  key: jwt key
// returns a claim if there was a success or an error otherwise
func Validate(token, key string) (*T1CGClaims, error) {
	parsed, err := jwt.ParseWithClaims(token, &T1CGClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method[%v]", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	return parsed.Claims.(*T1CGClaims), nil
}

// IssueAdminToken
// params
// 	userID: Unique ID that identifies the token bearer. Could be an email, mongo id, etc.
//  role: An integer that each application can represent however they choose. For example, an ecommerce site might look like
// 				1 = customer
//				2 = employee
//				3 = owner
// 	key: used to decode the signature of the token
//  issuer: who created the token
// 	expiry: when the token expires
// returns a signed token or an error.
func IssueAdminToken(userID, issuer, key string, role int64, expiry time.Time) (string, error) {
	claims := T1CGClaims{
		userID,
		role,
		jwt.StandardClaims{
			ExpiresAt: expiry.Unix(),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", fmt.Errorf("Failed to sign string")
	}

	return signedToken, nil
}
