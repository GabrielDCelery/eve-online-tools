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
	getMarketHistoryCmdRegionID uint64
	getMarketHistoryCmdTypeID   uint64
)

// getMarketHistoryCmd represents the getMarketHistory command
var getMarketHistoryCmd = &cobra.Command{
	Use:   "get-market-history",
	Short: "Retrieve market history for an item in a region",
	Long: `
Description:
  Retrieve market history for an item in a region

Examples(s):
  eve-online-tools api get-market-history --regionID=10000002 --typeID=32006
`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.GetMarketHistoryFromPublicApi(getMarketHistoryCmdRegionID, getMarketHistoryCmdTypeID)
	},
}

func init() {
	ApiCmd.AddCommand(getMarketHistoryCmd)

	getMarketHistoryCmd.Flags().Uint64Var(&getMarketHistoryCmdRegionID, "regionID", uint64(0), "Region ID")

	if err := getMarketHistoryCmd.MarkFlagRequired("regionID"); err != nil {
		log.Fatalln(err)
	}

	getMarketHistoryCmd.Flags().Uint64Var(&getMarketHistoryCmdTypeID, "typeID", uint64(0), "Type ID")

	if err := getMarketHistoryCmd.MarkFlagRequired("typeID"); err != nil {
		log.Fatalln(err)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getMarketHistoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getMarketHistoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
