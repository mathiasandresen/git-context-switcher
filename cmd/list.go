package cmd

import (
	"fmt"
	"gitc/config"

	"github.com/spf13/cobra"
)

func listContexts() error {
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	if len(cfg.Contexts) == 0 {
		fmt.Println("No contexts found. Use 'gitc init' to create one.")
		return nil
	}

	fmt.Println("Available contexts:")
	fmt.Println("------------------")
	for _, ctx := range cfg.Contexts {
		current := " "
		if ctx.Name == cfg.CurrentContext {
			current = "*"
		}
		fmt.Printf("%s %-20s (email: %s)\n", current, ctx.Name, ctx.Email)
	}
	return nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available Git contexts",
	Run: func(cmd *cobra.Command, args []string) {
		if err := listContexts(); err != nil {
			fmt.Printf("Error loading config: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
