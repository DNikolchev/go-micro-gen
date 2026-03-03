package cli

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add code components to an existing service",
	Long:  "Automatically scaffold handlers, services, and repositories into the current project.",
}
