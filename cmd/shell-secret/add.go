package main

import (
	"encoding/json"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:  "add [key] [value]",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]
		var input interface{}
		if err := json.Unmarshal([]byte(value), &input); err != nil {
			panic(err)
		}
		if err := shellSecret.Add(key, input); err != nil {
			panic(err)
		}
	},
}
