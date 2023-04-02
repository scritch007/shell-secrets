package main

import "github.com/spf13/cobra"

var deleteCmd = &cobra.Command{
	Use:  "delete [key]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := shellSecret.Delete(args[0]); err != nil {
			panic(err)
		}
	},
}
