package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type T1CGClaims struct {
	UserID string `json:"userId"`
	Role   int64  `json:"role"`
	jwt.StandardClaims
}
