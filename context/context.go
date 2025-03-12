package context

import (
	"fmt"
	"git-context-switcher/config"
	"git-context-switcher/utils"
	"os"
	"path/filepath"
)

func SwitchContext(ctx *config.Context) error {
	// Symlink private key
	err := utils.CreateSymlink(ctx.PrivateKey, filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa"))
	if err != nil {
		return err
	}

	// Symlink public key
	err = utils.CreateSymlink(ctx.PublicKey, filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa.pub"))
	if err != nil {
		return err
	}

	// Add the key to the agent
	err = utils.AddKey()
	if err != nil {
		return err
	}

	// Configure git email
	if ctx.Email != "" {
		err = utils.SetGitEmail(ctx.Email)
		if err != nil {
			return fmt.Errorf("failed to set git email: %v", err)
		}
	}

	fmt.Printf("Switched to context: %s\n", ctx.Name)
	if ctx.Email != "" {
		fmt.Printf("Git email set to: %s\n", ctx.Email)
	}
	return nil
}
