package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:     "gitc",
	Short:   "A tool to switch between Git contexts",
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		// If no arguments provided, show the list of contexts
		if len(args) == 0 {
			if err := listContexts(); err != nil {
				fmt.Printf("Error loading config: %v\n", err)
				os.Exit(1)
			}
		}
	},
}

func Execute() {
	// Check for any args before executing cobra
	if len(os.Args) > 1 {
		// Skip if the first argument is a known command or flag
		firstArg := os.Args[1]
		if firstArg != "init" && firstArg != "list" && firstArg != "switch" && !strings.HasPrefix(firstArg, "-") {
			// Insert "switch" as the first argument
			newArgs := make([]string, len(os.Args)+1)
			newArgs[0] = os.Args[0]
			newArgs[1] = "switch"
			copy(newArgs[2:], os.Args[1:])
			os.Args = newArgs
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
