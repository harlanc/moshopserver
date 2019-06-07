package services

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	jwt "github.com/dgrijalva/jwt-go"
)

var key = []byte("adfadf!@#2")

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

func getControllerAndAction(rawvalue string) (controller, action string) {
	vals := strings.Split(rawvalue, "/")
	return vals[2], vals[2] + "/" + vals[3]
}

var LoginUserId string

func FilterFunc(ctx *context.Context) {

	controller, action := getControllerAndAction(ctx.Request.RequestURI)
	token := ctx.Input.Header("x-nideshop-token")

	if action == "auth/loginByWeixin" {
		return
	}

	if token == "" {
		ctx.Redirect(401, "api/auth/loginByWeixin")
		return
	}
	LoginUserId = GetUserID(token)

	publiccontrollerlist := beego.AppConfig.String("controller::publicController")
	publicactionlist := beego.AppConfig.String("action::publicAction")

	if !strings.Contains(publiccontrollerlist, controller) && !strings.Contains(publicactionlist, action) {
		if LoginUserId == "" {
			ctx.Redirect(401, "/api/auth/loginByWeixin")
			//http.Redirect(ctx.ResponseWriter, ctx.Request, "/", http.StatusMovedPermanently)
		}
	}
}
