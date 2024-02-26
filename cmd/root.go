// Package cmd provides the entry point and command structure for the
// unix-utils application.
package cmd

import (
	"fmt"
	"os"

	_ "github.com/skraio/unix-utilities/cmd/ls" // Import ls command
	"github.com/skraio/unix-utilities/cmd/wc"   // Import wc command
	"github.com/spf13/cobra"                    // Import cobra library
)

// rootCmd represents the root command of the applicatoin.
var rootCmd = &cobra.Command{
	Use:   "unix-utils",
	Short: "Unix Utility Commands",
	Long:  "Implementation of various Unix utility commands in Go.",
	Run: func(cmd *cobra.Command, args []string) {
		// Default behavior to display usage when no command is specified.
		cmd.Usage()
	},
}

// init initializes the root command and adds subcommands to it.
func init() {
	rootCmd.AddCommand(wc.Cmd)
	// rootCmd.AddCommand(ls.lsCmd)
}

// Execute runs the root command, handling any errors.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
