package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	shellsecret "github.com/scritch007/shell-secrets"
	"github.com/spf13/cobra"
)

var shellSecret shellsecret.ShellSecret

func main() {
	rootCmd := cobra.Command{
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			shellSecret, err = shellsecret.New()
			if err != nil {
				if errors.Is(err, shellsecret.ErrEnvNotSetup) {
					if cmd.Name() == "setup" {
						return nil
					}
					if runtime.GOOS == "windows" {
						return fmt.Errorf(`run '%s setup'`, os.Args[0])
					}
					return fmt.Errorf(`run 'eval "$(%s setup)"'`, os.Args[0])
				}
				panic(err)
			}
			return nil
		},
	}
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(setupCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
}
