package main

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:  "get [key]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var in interface{}
		if err := shellSecret.Get(args[0], &in); err != nil {
			panic(err)
		}
		b, err := json.Marshal(&in)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", string(b))
	},
}
