// Package cmdflags provides structure for handling command-line flags.
package cmdflags

import (
	"os"

	"github.com/spf13/cobra"
)

// Flag represents a command-line flag with its properties.
type Flag struct {
	// Value indicates whether the flag is set.
	Value bool

	// Name is the full name of the flag.
	Name string

	// ShortHand is the shorthand abbreviation for the flag.
	ShortHand string

	// DefaultValue is the default value of the flag.
	DefaultValue bool

	// Description provides a brief description of the flag's purpose.
	Description string

	// Handler is the function that will be executed when the flag is encountered.
	Handler func(*os.File) int
}

func ParseFlags(flags []Flag, cmd *cobra.Command) {
	for i := range flags {
		f := &flags[i]
		cmd.Flags().BoolVarP(&f.Value, f.Name, f.ShortHand, f.DefaultValue, f.Description)
	}

	cmd.Flags().SetInterspersed(false)
}
