package main

import (
	"fmt"
	"github.com/lxn/walk"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var home string
var sessionDurationMinutes = int64(75)
var maxSessionDuration = sessionDurationMinutes * 60 // seconds
var username string
var warningInterval = int64(300) // 5 minutes (in seconds)

func init() {
	username = os.Getenv("USERNAME")
	home = os.Getenv("HOMEPATH")
}

func main() {
	notifyUser("Welcome, "+strings.Title(username)+"! You have "+strconv.FormatInt(sessionDurationMinutes, 10)+" minutes before your time is up.", "30")
	for duration := int64(0); duration < maxSessionDuration; duration++ {
		time.Sleep(1 * time.Second)

		remainder := (maxSessionDuration - duration) % warningInterval
		if remainder == 0 {
			timeleft := int64((maxSessionDuration - duration) / 60) // minutes
			msg := strconv.FormatInt(timeleft, 10) + " minutes left"
			go notifyUser(msg, "10")
		}
	}
	kickUser()
}

func kickUser() {
	notifyUser("Logging you off in 30 seconds!!!", "20")
	time.Sleep(30 * time.Second)
	shutDown()
}

func notifyUser(msg string, timeout string) {
	walk.MsgBox(nil, "Tardis Update", msg, walk.MsgBoxIconWarning)
	/*
		cmd := exec.Command("msg", username, "/TIME:"+timeout, msg)
		err := cmd.Start()
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		cmd.Wait()*/
}

func shutDown() {
	cmd := exec.Command("shutdown", "/l")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	cmd.Wait()
}
