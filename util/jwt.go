package util

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateJWT(id string) (string, error) {
	// claims/payload data
	claims := jwt.MapClaims{
		"id": id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)                    // algorithm/token type
	tokenString, err := token.SignedString([]byte(viper.GetString("SECRET_KEY"))) // verify signature
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpeted signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}
