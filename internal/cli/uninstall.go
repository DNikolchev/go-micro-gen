package cli

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall go-micro-gen from your system",
	Run: func(cmd *cobra.Command, args []string) {
		execPath, err := os.Executable()
		if err != nil {
			color.Red("❌ Could not determine executable path: %v", err)
			return
		}

		color.Yellow("Uninstalling go-micro-gen from: %s", execPath)

		err = os.Remove(execPath)
		if err != nil {
			color.Red("❌ Failed to uninstall: %v", err)
			color.Cyan("\nYou may need to manually delete it or run the uninstall command with sudo/administrator privileges.")
			return
		}

		color.Green("✅ go-micro-gen successfully uninstalled!")
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
