package service

import (
	"github.com/growmax/noti/db"
	"github.com/growmax/noti/model"
)

func SendNotification(noti model.Notification) {

	client := GetCentriConnection()
	Centripush(noti.User, noti.Notification, *client)
	recieveroption, _ := db.GetRecieverOption(noti.User)
	if recieveroption.Webpush {
		browsers, err := db.GetBrowser(noti.User, 1) //last 3 browsers
		if err == nil {
			for _, browser := range browsers {
				go SendWebpush(browser.Browser, noti.Notification)
			}
		}
	}
}
