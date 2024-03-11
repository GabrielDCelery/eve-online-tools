/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/GabrielDCelery/eve-online-tools-cli/cmd/api"
	"github.com/GabrielDCelery/eve-online-tools-cli/cmd/transforms"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eve-online-tools",
	Short: "CLI commands for various EVE Online related operations",
	Long: `
Description:
  CLI commands for various EVE Online related operations

Example(s):
  eve-online-tools api get-market-orders --regionID 10000002`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.eve-online-tools.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(api.ApiCmd)
	rootCmd.AddCommand(transforms.TransformsCmd)

	// rootCmd.AddCommand(filter.FilterCmd)
	// rootCmd.AddCommand(calculate.CalculateCmd)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
