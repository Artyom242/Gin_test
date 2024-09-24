package utils

import "github.com/dgrijalva/jwt-go"

func GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
