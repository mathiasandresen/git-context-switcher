package utils

import (
	"os/exec"
)

func SetGitEmail(email string) error {
	cmd := exec.Command("git", "config", "--global", "user.email", email)
	return cmd.Run()
}
