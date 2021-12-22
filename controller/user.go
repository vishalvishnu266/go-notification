package controller

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/growmax/noti/db"
	"github.com/growmax/noti/model"
)

func CreateUser(c *gin.Context) {
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		c.String(500, "jwt error")
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.String(500, "jwt error")
		return
	}
	isuserexist, err := db.IsUserExist(claims["sub"].(string))
	if err != nil {
		c.String(500, "adding user error")
		return
	}
	if !isuserexist {

		var User model.User
		User.User = claims["sub"].(string)
		User.Tenantid = claims["tenantId"].(string)
		User.IsSeller = claims["isSeller"].(bool)
		User.PriEmail = claims["email"].(string)
		User.SecEmail = claims["secondaryEmail"].(string)
		User.PriPhone = claims["phoneNumber"].(string)
		User.SecPhone = claims["secondaryPhoneNumber"].(string)
		User.CreatedAt = time.Now()
		err = db.CreateUser(User)
		if err != nil {
			c.String(500, "mongo error")
			return
		}
		_ = db.AddRecieverOption(model.RecieverOption{
			User:    claims["sub"].(string),
			Webpush: false,
		})
	}
	c.String(200, "success")

}
