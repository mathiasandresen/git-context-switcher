package cmd

import (
	"fmt"
	"gitc/config"
	"gitc/context"
	"os"

	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch [context]",
	Short: "Switch to a different Git context",
	Run:   switchContext,
}

func init() {
	rootCmd.AddCommand(switchCmd)
}

func switchContext(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: Context name is required")
		os.Exit(1)
	}

	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No configuration file found. Run 'gitc init' to create one.")
			os.Exit(1)
		}
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Find the context in the config
	contextName := args[0]
	var contextToSwitch *config.Context
	for _, ctx := range cfg.Contexts {
		if ctx.Name == contextName {
			contextToSwitch = &ctx
			break
		}
	}

	if contextToSwitch == nil {
		fmt.Printf("Context '%s' not found\n", contextName)
		os.Exit(1)
	}

	// Switch context by creating symlinks
	err = context.SwitchContext(contextToSwitch)
	if err != nil {
		fmt.Printf("Error switching context: %v\n", err)
		os.Exit(1)
	}
}
