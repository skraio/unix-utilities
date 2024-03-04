// Package ls provides functionality for listing files in a directory with
// various formatting options.
package ls

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/skraio/unix-utilities/cmdflags"
	"github.com/spf13/cobra"
)

// Terminal color codes.
const (
	Reset = "\033[0m"
	Blue  = "\033[34m"
)

// FileAttributes holds information about a file.
type FileAttributes struct {
	fileMode os.FileMode
	ulink    string
	uid      string
	fileSize string
	modTime  string
}

// OutputEntry represents a file entry with its attributes.
type OutputEntry struct {
	fileName       string
	fileAttributes FileAttributes
}

// lsFlags holds flags for ls command.
type lsFlags struct {
	longForm    bool
	all         bool
	readable    bool
	timeSort    bool
	reverseSort bool
}

var pFlags lsFlags

// flags definition for ls command.
var flags = []cmdflags.Flag{
	{Value: &pFlags.longForm, Name: "long", ShortHand: "l", DefaultValue: false, Description: "detailed file information display"},
	{Value: &pFlags.all, Name: "all", ShortHand: "a", DefaultValue: false, Description: "show all files, including hidden ones"},
	{Value: &pFlags.readable, Name: "readableSize", ShortHand: "h", DefaultValue: false, Description: "human-readable size format"},
	{Value: &pFlags.timeSort, Name: "sort", ShortHand: "t", DefaultValue: false, Description: "sort output by modification time"},
	{Value: &pFlags.readable, Name: "reverse", ShortHand: "r", DefaultValue: false, Description: "reverse output order"},
}

// Cmd represents the 'ls' command configuration using Cobra.
var Cmd = &cobra.Command{
	Use:   "ls [-f flags]",
	Short: "List directory content with optional formatting flags.",
	Run: func(cmd *cobra.Command, args []string) {
		err := executeLs(args)
		if err != nil {
			log.Print(err.Error())
			return
		}
	},
}

// init initializes the 'ls' command by setting up flags.
func init() {
	cmdflags.ParseFlags(flags, Cmd)
	Cmd.PersistentFlags().BoolP("help", "", false, "help for this command")
}

// executeLs executes the ls command with given arguments.
func executeLs(args []string) error {
	n := len(args)
	if n == 0 {
		curDir := "."
		list, _ := execute(curDir)
		printList(list, curDir, n)
	} else {
		for _, arg := range args {
			list, err := execute(arg)
			if err != nil {
				if os.IsNotExist(err) {
					continue
				}
				return err
			}

			printList(list, arg, n)
		}
	}

	return nil
}

// execute executes ls command in the given directory.
func execute(dir string) ([]OutputEntry, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("Directory", dir, "does not exists")
		return nil, err
	}

	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	content, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}

	sortByName(content)

	if pFlags.timeSort {
		sortByModTime(content)
	}
	if pFlags.reverseSort {
		reverseOrder(content)
	}

	output := []OutputEntry{}
	for _, file := range content {
		if !pFlags.all && strings.HasPrefix(file.Name(), ".") {
			continue
		}

		entry := OutputEntry{}
		entry.fileName = colorize(file)

		if pFlags.longForm {
			entry.fileAttributes, err = longFormat(file)
			if err != nil {
				return nil, err
			}
		}
		output = append(output, entry)
	}

	return output, nil
}

// printList prints the ls output.
func printList(output []OutputEntry, dir string, n int) {
	if n > 1 {
		fmt.Printf("%s:\n", dir)
	}

	if !pFlags.longForm {
		for _, o := range output {
			fmt.Printf("%s ", o.fileName)
		}
		fmt.Println()
		if n > 1 {
			fmt.Println()
		}
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
	defer w.Flush()

	for _, o := range output {
		a := o.fileAttributes
		fmt.Fprintf(w, "%s\t\t%s\t\t%s\t\t%s\t\t%s\t\t%s",
			a.fileMode, a.ulink, a.uid, a.fileSize, a.modTime, o.fileName)
		fmt.Fprintln(w)
	}

	if n > 1 {
		fmt.Println()
	}
}

// sortByName sorts files by name.
func sortByName(content []os.FileInfo) {
	sort.Slice(content, func(i, j int) bool {
		if content[i].IsDir() && !content[j].IsDir() {
			return true
		}
		if !content[i].IsDir() && content[j].IsDir() {
			return false
		}
		if content[i].Name()[0] == '.' && content[j].Name()[0] != '.' {
			return true
		}
		if content[i].Name()[0] != '.' && content[j].Name()[0] == '.' {
			return false
		}
		return content[i].Name() < content[j].Name()
	})
}

// colorize applies color to directory names.
func colorize(file fs.FileInfo) string {
	if file.IsDir() {
		return Blue + file.Name() + Reset
	}
	return file.Name()
}
