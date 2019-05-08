package services

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

const secret string = "adf1212!@$"

func GetUserID(token string) {

}

func Parse(tokenstr string) {

	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if token.Valid {
		fmt.Println("You look nice today")
		//return token.Claims
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

}

func create(userInfo string) string {

	token, err := jwt.SigningMethodHS256.Sign(userInfo, secret)
	if err == nil {
		return token
	}
	return ""
}
