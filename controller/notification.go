package controller

import (
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/growmax/noti/db"
)

type Pageno struct {
	Page string `form:"page"`
}

func GetNotification(c *gin.Context) {
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	//#####################3
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
	var p Pageno
	if c.ShouldBindQuery(&p) != nil {
		c.String(http.StatusBadRequest, "invalid parameter")
		return
	}
	i, err := strconv.Atoi(p.Page)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid parameter")
		return
	}
	notifications, err := db.GetNotification(claims["sub"].(string), i)
	if err != nil {
		c.String(500, "mongo error")
		return
	}
	c.JSON(200, notifications)

}
