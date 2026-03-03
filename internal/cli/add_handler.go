package cli

import (
	"fmt"

	"github.com/Aro-M/go-micro-gen/internal/generator"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	flagHandlerName  string
	flagHandlerRoute string
)

var addHandlerCmd = &cobra.Command{
	Use:   "handler",
	Short: "Add a new REST handler component",
	Long:  "Scaffold a fully compliant chi-based router HTTP handler inside 'internal/transport/httpx'.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if flagHandlerName == "" {
			return fmt.Errorf("--name is required")
		}
		if flagHandlerRoute == "" {
			return fmt.Errorf("--route is required")
		}

		color.Cyan("\n🚀 Adding handler '%s' mapped to route '%s' ...\n", flagHandlerName, flagHandlerRoute)

		err := generator.AddHandler(flagHandlerName, flagHandlerRoute)
		if err != nil {
			color.Red("❌ Add handler failed: %v", err)
			return err
		}

		color.Green("✅ Handler generated successfully!")
		return nil
	},
}

func init() {
	addHandlerCmd.Flags().StringVar(&flagHandlerName, "name", "", "Name of the handler (e.g. User)")
	addHandlerCmd.Flags().StringVar(&flagHandlerRoute, "route", "", "Base route for the handler (e.g. /users)")
	addCmd.AddCommand(addHandlerCmd)
}
