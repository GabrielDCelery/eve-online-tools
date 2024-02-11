/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package filter

import (
	"log"

	"github.com/GabrielDCelery/eve-online-tools-cli/lib"
	"github.com/spf13/cobra"
)

var (
	marketOrdersCmdTypeID uint64
)

// marketOrdersCmd represents the marketOrders command
var marketOrdersCmd = &cobra.Command{
	Use:   "market-orders",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.FilterMarketOrders(marketOrdersCmdTypeID)
	},
}

func init() {
	FilterCmd.AddCommand(marketOrdersCmd)

	marketOrdersCmd.Flags().Uint64Var(&marketOrdersCmdTypeID, "typeID", uint64(0), "Type ID")

	if err := marketOrdersCmd.MarkFlagRequired("typeID"); err != nil {
		log.Fatalln(err)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// marketOrdersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// marketOrdersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
