package utils

import (
	"fmt"
	"os"
)

func CreateSymlink(target string, link string) error {
	// First, remove the existing symlink or file if any
	if _, err := os.Stat(link); err == nil {
		err := os.Remove(link)
		if err != nil {
			return fmt.Errorf("failed to remove existing symlink: %v", err)
		}
	}

	// Delete the old symlink if it exists
	if _, err := os.Lstat(link); err == nil {
		err := os.Remove(link)
		if err != nil {
			return fmt.Errorf("failed to remove existing symlink: %v", err)
		}
	}

	// Create the symlink
	err := os.Symlink(target, link)
	if err != nil {
		return fmt.Errorf("failed to create symlink: %v", err)
	}

	return nil
}
