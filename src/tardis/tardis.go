package main

import (
	"flag"
	"fmt"
	"strings"
	"tardis/logout"
	"tardis/popup"
	"tardis/userinfo"
	"time"
)

var sessionDurationMinutes = flag.Int("t", 75, "maximum session time in minutes")
var home string
var maxSessionDuration int
var username string
var warningInterval = 300 // 5 minutes (in seconds)

func init() {
	home = userinfo.GetUserHome()
	username = userinfo.GetUserLogin()
}

func main() {
	flag.Parse()

	maxSessionDuration = *sessionDurationMinutes * 60 // seconds
	time.Sleep(5 * time.Second)
	msg := fmt.Sprintf("Welcome %s! You have %v minutes before your time is up.", strings.Title(username), *sessionDurationMinutes)
	go notifyUser(msg)
	for duration := 0; duration < maxSessionDuration; duration++ {
		time.Sleep(1 * time.Second)

		remainder := (maxSessionDuration - duration) % warningInterval
		if remainder == 0 {
			timeleft := (maxSessionDuration - duration) / 60 // minutes
			msg = fmt.Sprintf("%v minutes left", timeleft)
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
