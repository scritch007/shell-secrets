package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		keys, err := shellSecret.List()
		if err != nil {
			panic(err)
		}
		fmt.Printf("available keys:\n")
		for _, k := range keys {
			fmt.Printf("  - %s\n", k)
		}
	},
}
