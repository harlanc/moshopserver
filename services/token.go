package services

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

const key string = "adf1212!@$"

type CustomClaims struct {
	UserID string `json:"userid"`
	jwt.StandardClaims
}

func GetUserID(tokenstr string) string {

	token := Parse(tokenstr)

	if token == nil {
		return ""
	}

	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims.UserID
	}

	return ""
}

func Parse(tokenstr string) *jwt.Token {

	token, err := jwt.ParseWithClaims(tokenstr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if token.Valid {
		return token
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
	return nil

}

func Create(userid string) string {
	claims := CustomClaims{
		userid, jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := token.SignedString(key)

	if err == nil {
		return tokenstr
	}
	return ""
}

func Verify(tokenstr string) bool {

	token := Parse(tokenstr)
	return token != nil

}
