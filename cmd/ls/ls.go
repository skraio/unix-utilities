package ls

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/skraio/unix-utilities/cmdflags"
	"github.com/spf13/cobra"
)

type FileAttributes struct {
	fileMode os.FileMode
	ulink    string
	uid      string
	fileSize string
	modTime  string
}

type OutputEntry struct {
	fileName       string
	fileAttributes FileAttributes
}

var flags = []cmdflags.Flag{
	{Value: false, Name: "long", ShortHand: "l", DefaultValue: false, Description: "detailed file information display", Handler: nil},
	{Value: false, Name: "all", ShortHand: "a", DefaultValue: false, Description: "show all files, including hidden ones", Handler: nil},
	{Value: false, Name: "readableSize", ShortHand: "h", DefaultValue: false, Description: "human-readable size format", Handler: nil},
	{Value: false, Name: "sort", ShortHand: "t", DefaultValue: false, Description: "sort output by modification time", Handler: nil},
	{Value: false, Name: "reverse", ShortHand: "r", DefaultValue: false, Description: "reverse output order", Handler: nil},
}

var Cmd = &cobra.Command{
	Use:   "ls [-f flags]",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		list, fLongForm := executeLs(args)
		printList(list, fLongForm)
	},
}

func init() {
	cmdflags.ParseFlags(flags, Cmd)
	Cmd.PersistentFlags().BoolP("help", "", false, "help for this command")
}

func executeLs(args []string) ([]OutputEntry, bool) {
	fLongForm := false
	fAll := false
	fHuman := false
	fTimeSort := false
	fReverseSort := false
	for _, f := range flags {
		if !f.Value {
			continue
		}

		switch f.ShortHand {
		case "l":
			fLongForm = true
		case "a":
			fAll = true
		case "h":
			fHuman = true
		case "t":
			fTimeSort = true
		case "r":
			fReverseSort = true
		}
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Print(err.Error())
	}

	f, err := os.Open(dir)
	if err != nil {
		log.Print(err.Error())
	}
	defer f.Close()

	content, err := f.Readdir(-1)
	if err != nil {
		log.Print(err.Error())
	}

	sortByName(content)

	if fTimeSort {
		sortByModTime(content)
	}
	if fReverseSort {
		reverseOrder(content)
	}

	output := []OutputEntry{}
	for _, file := range content {
		newEntry := &OutputEntry{}

		if !fAll && file.Name()[0] == '.' {
			continue
		}

		newEntry.fileName = colorize(file)

		if !fLongForm {
			output = append(output, *newEntry)
			continue
		}

		newEntry.fileAttributes = longFormat(file, fHuman)
		output = append(output, *newEntry)
	}

	return output, fLongForm
}

func printList(output []OutputEntry, fLongForm bool) {
	if !fLongForm {
		for _, o := range output {
			fmt.Print(o.fileName, " ")
		}
		fmt.Println()
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
	for _, o := range output {
		a := o.fileAttributes
		fmt.Fprintf(w, "%s\t %2s\t %s\t %4s\t %s\t %s",
			a.fileMode, a.ulink, a.uid, a.fileSize, a.modTime, o.fileName)
		fmt.Fprintln(w)
	}

	w.Flush()
}

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

func colorize(file fs.FileInfo) string {
	const (
		Reset = "\033[0m"
		Blue  = "\033[34m"
	)

	if file.IsDir() {
		return Blue + file.Name() + Reset
	}
	return file.Name()
}
