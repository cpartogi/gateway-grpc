package helper

import (
	"os"

	"github.com/golang-jwt/jwt"
)

func GetDataFromToken(token string) (userId string, err error) {

	resToken, err := ParseToken(token)
	if err != nil {
		return
	}

	claims := resToken.Claims.(jwt.MapClaims)

	userId = claims["id"].(string)

	return
}

func ParseToken(tokenString string) (*jwt.Token, error) {

	tokenKey := "TOKEN_KEY"

	tokenKeyString := os.Getenv(tokenKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, jwt.ErrInvalidKeyType
		}

		return []byte(tokenKeyString), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	return token, err
}
