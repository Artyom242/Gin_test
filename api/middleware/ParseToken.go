package middleware

import "github.com/dgrijalva/jwt-go"

func GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
