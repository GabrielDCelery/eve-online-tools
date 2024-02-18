/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package api

import (
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "Commands to access the public EVE Online API",
	Long: `
Description:
  Commands to access the public EVE Online API

Example(s):
  eve-online-tools api get-market-orders --regionID 10000002
  eve-online-tools api get-market-history --regionID=10000002 --typeID=32006
`,
	/*
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("api called")
		},
	*/
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
