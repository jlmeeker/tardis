package logout

import (
	"fmt"
	"os/exec"
)

func LogoutUser(username string) {
	cmd := exec.Command("shutdown", "/l")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	cmd.Wait()
}
