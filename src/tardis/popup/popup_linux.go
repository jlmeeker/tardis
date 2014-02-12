package popup

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

var progs = []string{"zenity"}

func ShowPopUpWindow(msg string) {
	var err error
	for _, prog := range progs {
		if prog == "zenity" {
			err = useZenity(msg)
		} else {
			err = errors.New("invalid popup app")
		}

		if err == nil {
			break
		}
	}

	if err != nil {
		fmt.Printf("popup error: %v\n", err)
		os.Exit(1)
	}
}

func useZenity(msg string) (err error) {
	cmd := exec.Command("zenity", "--warning", "--title", "Tardis Update", "--text", msg)
	err = cmd.Start()
	if err == nil {
		cmd.Wait()
	}

	return
}
