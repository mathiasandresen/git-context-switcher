package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func AddKey() error {
	// Call ssh-add to add the key to the agent
	cmd := exec.Command("ssh-add", filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa"))
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to add key to agent: %v", err)
	}

	return nil
}
