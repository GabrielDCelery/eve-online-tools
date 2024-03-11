/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transforms

import (
	"log"

	"github.com/GabrielDCelery/eve-online-tools-cli/lib"
	"github.com/spf13/cobra"
)

var (
	transformMarketOrdersCmdFrom string
	transformMarketOrdersCmdTo   string
)

// transformMarketOrdersCmd represents the transformMarketOrders command
var transformMarketOrdersCmd = &cobra.Command{
	Use:   "market-orders",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.TransformMarketOrdersConfig{
			From: transformMarketOrdersCmdFrom,
			To:   transformMarketOrdersCmdTo,
		}
		lib.TransforMarketOrders(&config)
	},
}

func init() {
	TransformsCmd.AddCommand(transformMarketOrdersCmd)

	transformMarketOrdersCmd.Flags().StringVar(&transformMarketOrdersCmdFrom, "from", "", "From")

	if err := transformMarketOrdersCmd.MarkFlagRequired("from"); err != nil {
		log.Fatalln(err)
	}

	transformMarketOrdersCmd.Flags().StringVar(&transformMarketOrdersCmdTo, "to", "", "To")

	if err := transformMarketOrdersCmd.MarkFlagRequired("to"); err != nil {
		log.Fatalln(err)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transformMarketOrdersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transformMarketOrdersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
