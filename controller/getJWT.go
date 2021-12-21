package controller

import (
	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"
	"github.com/growmax/noti/service"
)

func GetJwt(c *gin.Context) {
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
	t, err := service.GenerateJwt(claims["sub"].(string))
	if err != nil {
		c.String(500, "Token generation error")
		return
	}
	c.String(200, t)

}
