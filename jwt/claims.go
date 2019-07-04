package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type T1CGClaims struct {
	UserID string `json:"userId"`
	Role   Role   `json:"role"`
	jwt.StandardClaims
}

type TokenRole struct {
	Token        string `json:"token"`
	RoleRequired Role   `json:"roleRequired`
}

type Role int64

const (
	None Role = iota
	User
	System
	SuperUser
	SiteAdmin
	Admin
)
