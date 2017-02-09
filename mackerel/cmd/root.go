package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "graph-annotations",
	Short: "Graph-annotations is a very simple tool for mackerel graph annotations API.",
	Long:  "Complete API documentation is available at https://mackerel.io/api-docs/entry/graph-annotations",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cobra.OnInitialize()
}
