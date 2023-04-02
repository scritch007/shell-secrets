package main

import (
	shellsecret "github.com/scritch007/shell-secrets"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use: "setup",
	Run: func(cmd *cobra.Command, args []string) {
		if err := shellsecret.Setup(); err != nil {
			panic(err)
		}
	},
}
