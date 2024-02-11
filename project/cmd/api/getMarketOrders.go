/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package api

import (
	"log"

	"github.com/GabrielDCelery/eve-online-tools-cli/lib"

	"github.com/spf13/cobra"
)

var (
	getMarketOrdersCmdRegionID uint64
)

// getMarketOrdersCmd represents the getMarketOrders command
var getMarketOrdersCmd = &cobra.Command{
	Use:   "get-market-orders",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.GetMarketOrdersFromPublicApiConfig{
			RegionID: getMarketOrdersCmdRegionID,
		}

		lib.GetMarketOrdersFromPublicApi(&config)
	},
}

func init() {
	ApiCmd.AddCommand(getMarketOrdersCmd)

	getMarketOrdersCmd.Flags().Uint64Var(&getMarketOrdersCmdRegionID, "regionID", uint64(0), "Region ID")

	if err := getMarketOrdersCmd.MarkFlagRequired("regionID"); err != nil {
		log.Fatalln(err)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getMarketOrdersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getMarketOrdersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
