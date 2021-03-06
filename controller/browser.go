package controller

import (
	"encoding/json"
	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/growmax/noti/db"
	"github.com/growmax/noti/model"
	"strings"
	"time"
)

type Browser struct {
	BrowserEnd string `form:"browser"`
}

func AddBrowser(c *gin.Context) {
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
	broAdd.Browser = b.BrowserEnd
	broAdd.CreatedAt = time.Now()
	err = db.AddBrowser(broAdd)
	if err != nil {
		c.String(500, "internal error")
		return
	}
	c.String(200, "Success")
}
