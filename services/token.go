package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/harlanc/moshopserver/utils"
)

var key = []byte("adfadf!@#2")

var expireTime = 20

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
			fmt.Println("The token is expired or not valid.")
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
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(expireTime)).Unix(),
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
		data := utils.GetHTTPRtnJsonData(401, "need relogin")
		ctx.Output.JSON(data, true, false)
		ctx.Redirect(200, "/")
		return
	}
	LoginUserId = GetUserID(token)

	publiccontrollerlist := beego.AppConfig.String("controller::publicController")
	publicactionlist := beego.AppConfig.String("action::publicAction")

	if !strings.Contains(publiccontrollerlist, controller) && !strings.Contains(publicactionlist, action) {
		if LoginUserId == "" {
			data := utils.GetHTTPRtnJsonData(401, "need relogin")
			ctx.Output.JSON(data, true, false)
			ctx.Redirect(200, "/")
			//http.Redirect(ctx.ResponseWriter, ctx.Request, "/", http.StatusMovedPermanently)
		}
	}
}
