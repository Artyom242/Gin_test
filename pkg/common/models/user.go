package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Password  string `json:"password"`
	AuthToken string `json:"auth_token"`
	jwt.StandardClaims
}
