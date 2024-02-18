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
	Short: "Retrieve market orders in a region",
	Long: `
Description:
  Retrieve market orders in a region

Examples(s):
  eve-online-tools api get-market-orders --regionID 10000002
`,
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
