package logout

import (
	"fmt"
	"os/exec"
)

func LogoutUser(username string) {
	cmd := exec.Command("pkill", "-u", username)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	cmd.Wait()
}
