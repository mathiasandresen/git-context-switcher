package cmd

import (
	"fmt"
	"gitc/config"
	"gitc/context"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitc",
	Short: "A tool to switch between Git contexts",
	Run: func(cmd *cobra.Command, args []string) {
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

		// Determine the context to switch to (either provided or current)
		contextName := cfg.CurrentContext
		if len(args) > 0 {
			contextName = args[0]
		}

		// Find the context in the config
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
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
