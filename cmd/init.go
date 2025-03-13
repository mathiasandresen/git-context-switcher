package cmd

import (
	"fmt"
	"gitc/config"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize gitc",
	Long: `Initialize gitc by creating a config file with a default context.
The config file will be created at ~/.git-contexts.yaml if it doesn't exist.`,
	Run: func(cmd *cobra.Command, args []string) {
		configFile := filepath.Join(os.Getenv("HOME"), ".git-contexts.yaml")

		// Check if config file already exists
		if _, err := os.Stat(configFile); err == nil {
			fmt.Println("Configuration file already exists at:", configFile)
			return
		}

		// Create default configuration
		defaultConfig := &config.Config{
			CurrentContext: "default",
			Contexts: []config.Context{
				{
					Name:       "default",
					PrivateKey: filepath.Join(os.Getenv("HOME"), ".ssh/id_rsa"),
					PublicKey:  filepath.Join(os.Getenv("HOME"), ".ssh/id_rsa.pub"),
					Email:      "",
				},
			},
		}

		// Marshal config to YAML
		yamlData, err := yaml.Marshal(defaultConfig)
		if err != nil {
			fmt.Printf("Error creating config: %v\n", err)
			os.Exit(1)
		}

		// Write config file
		err = os.WriteFile(configFile, yamlData, 0600)
		if err != nil {
			fmt.Printf("Error writing config file: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Initialized gitc")
		fmt.Println("Config file created at:", configFile)
		fmt.Println("\nNext steps:")
		fmt.Println("1. Set your email in the config file")
		fmt.Println("2. Add additional contexts as needed")
		fmt.Printf("\nExample context:\n" +
			"contexts:\n" +
			"  - name: work\n" +
			"    private_key: ~/.ssh/work_id_rsa\n" +
			"    public_key: ~/.ssh/work_id_rsa.pub\n" +
			"    email: your@work-email.com\n")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
