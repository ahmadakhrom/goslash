package models

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type Login struct {
	Username  string    `gorm:"not null;size:30" validate:"required" json:"username"`
	Password  string    `gorm:"not null;size:100" validate:"required" json:"password"`
}
