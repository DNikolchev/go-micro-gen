package cli

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var banner = `
  __  __  ____    ____   _____  _   _
 |  \/  |/ ___|  / ___| | ____|| \ | |
 | |\/| |\___ \ | |  _  |  _|  |  \| |
 | |  | | ___) || |_| | | |___ | |\  |
 |_|  |_||____/  \____| |_____||_| \_|

 The Ultimate Microservice Scaffolder
`

var rootCmd = &cobra.Command{
	Use:   "go-micro-gen",
	Short: "Production-ready Go microservice generator",
	Long:  color.CyanString(banner) + "\n  Generate fully-wired, production-ready Go microservices in seconds.",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(addCmd)
}
