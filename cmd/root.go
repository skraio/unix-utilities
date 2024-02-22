package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	options string
	rootCmd = &cobra.Command{}
)

func init() {
	rootCmd.AddCommand(wcCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
