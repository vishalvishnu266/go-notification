package controller

import (
	"encoding/json"
	"log"
	"strings"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/growmax/noti/db"
	"github.com/growmax/noti/model"
)

type Browser struct {
	// User       string `form:"user"`
	BrowserEnd string `form:"browser"`
}

func AddBrowser(c *gin.Context) {
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	//#####################3
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		c.String(500, "internal error")
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		log.Println(claims["sub"])
	} else {
		c.String(500, "internal error")
		return
	}
	//##############################################
	var b Browser
	s := webpush.Subscription{}
	if c.ShouldBindQuery(&b) == nil {
		err := json.Unmarshal([]byte(b.BrowserEnd), &s)
		if err != nil {
			c.String(400, "invalid data")
			return
		}
	}
	var broAdd model.Browser
	broAdd.User = claims["sub"].(string)
	broAdd.Browser = s
	err = db.AddBrowser(broAdd)
	if err != nil {
		c.String(500, "internal error")
		return
	}
	c.String(200, "Success")
}
