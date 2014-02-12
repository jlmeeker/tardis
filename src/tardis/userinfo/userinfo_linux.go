package userinfo

import (
	"os"
)

func GetUserLogin() (login string) {
	login = os.Getenv("USER")
	return
}

func GetUserHome() (home string) {
	home = os.Getenv("HOME")
	return
}
