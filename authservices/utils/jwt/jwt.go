package jwt

import (
	"fmt"
	"time"

	// "time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(details map[string]string) (string, error) {
	mySigningKey := []byte("elvisSecreyKey")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["firstname"] = details["firstname"]
	claims["lastname"] = details["lastname"]
	claims["email"] = details["email"]
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func DecodeJWT(tokenBearer string) error {
	var tokenString = tokenBearer

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("elvisSecreyKey"), nil
	})

	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["foo"], claims["firstname"])
			return nil
		} else {
			return err
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return err
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return err
		} else {
			return err
		}
	} else {
		return err
	}
}
