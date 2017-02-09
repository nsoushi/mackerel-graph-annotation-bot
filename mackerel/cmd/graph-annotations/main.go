package main

import (
	"fmt"
	"github.com/nsoushi/mackerel-graph-annotation-bot/mackerel/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
