package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/growmax/noti/db"
	"github.com/growmax/noti/model"
	"strings"
)

func GetRecieverOption(c *gin.Context) {
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		c.String(500, "internal error")
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.String(500, "internal error")
		return
	}
	db.AddRecieverOption(model.RecieverOption{
		User:    claims["sub"].(string),
		Webpush: false,
	})
	option, err := db.GetRecieverOption(claims["sub"].(string))
	if err != nil {
		c.String(500, "mongo error")
		return
	}
	c.JSON(200, option)

}
func EnableWebpush(c *gin.Context) {
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		c.String(500, "internal error")
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.String(500, "internal error")
		return
	}
	err = db.UpdateWebPush(claims["sub"].(string))
	if err != nil {
		c.String(500, "mongo error")
		return
	}
	c.String(200, "success")

}

func EnableSms(c *gin.Context) {
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		c.String(500, "internal error")
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.String(500, "internal error")
		return
	}
	err = db.UpdateSms(claims["sub"].(string))
	if err != nil {
		c.String(500, "mongo error")
		return
	}
	c.String(200, "success")

}

func EnableEmail(c *gin.Context) {
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		c.String(500, "internal error")
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.String(500, "internal error")
		return
	}
	err = db.UpdateEmail(claims["sub"].(string))
	if err != nil {
		c.String(500, "mongo error")
		return
	}
	c.String(200, "success")

}
