package main

import (
	"strconv"
	"strings"
	"tardis/logout"
	"tardis/popup"
	"tardis/userinfo"
	"time"
)

var home string
var sessionDurationMinutes = int64(75)
var maxSessionDuration = sessionDurationMinutes * 60 // seconds
var username string
var warningInterval = int64(300) // 5 minutes (in seconds)

func init() {
	home = userinfo.GetUserHome()
	username = userinfo.GetUserLogin()
}

func main() {
	time.Sleep(5 * time.Second)
	go notifyUser("Welcome, " + strings.Title(username) + "! You have " + strconv.FormatInt(sessionDurationMinutes, 10) + " minutes before your time is up.")
	for duration := int64(0); duration < maxSessionDuration; duration++ {
		time.Sleep(1 * time.Second)

		remainder := (maxSessionDuration - duration) % warningInterval
		if remainder == 0 {
			timeleft := int64((maxSessionDuration - duration) / 60) // minutes
			msg := strconv.FormatInt(timeleft, 10) + " minutes left"
			go notifyUser(msg)
		}
	}
	kickUser()
}

func kickUser() {
	go notifyUser("Logging you off in 30 seconds!!!")
	time.Sleep(30 * time.Second)
	logout.LogoutUser(username)
}

func notifyUser(msg string) {
	popup.ShowPopUpWindow(msg)
}
