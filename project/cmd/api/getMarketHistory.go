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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
