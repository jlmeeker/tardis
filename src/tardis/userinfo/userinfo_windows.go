package userinfo

import (
	"os"
)

func GetUserLogin() (login string) {
	login = os.Getenv("USERNAME")
	return
}

func GetUserHome() (home string) {
	home = os.Getenv("HOMEPATH")
	return
}
