// Package wc provides functionality for counting lines, words, bytes and
// longest line length in files.
package wc

import (
	"fmt"
	"io"
	"log"
	"os"
	"text/tabwriter"

	"github.com/skraio/unix-utilities/cmdflags"
	"github.com/spf13/cobra"
)

// flags represents the command-line flags to control the behavior of the 'wc' command.
var flags = []cmdflags.Flag{
	{Value: false, Name: "lines", ShortHand: "l", DefaultValue: false, Description: "print the newline counts", Handler: lineCounter},
	{Value: false, Name: "words", ShortHand: "w", DefaultValue: false, Description: "print the word counts", Handler: wordCounter},
	{Value: false, Name: "bytes", ShortHand: "c", DefaultValue: false, Description: "print the byte counts", Handler: byteCounter},
	{Value: false, Name: "longest", ShortHand: "L", DefaultValue: false, Description: "print the length of the longest line", Handler: longestLine},
}

// Cmd represents the 'wc' command configuration using Cobra.
var Cmd = &cobra.Command{
	Use:   "wc [file]... [-f flags]",
	Short: "Line, word, byte and longest line count",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stats, longestLine := executeWc(args)
		printStats(args, stats, longestLine)
	},
}

// init initializes the 'wc' command by setting up flags.
func init() {
	for i := range flags {
		f := &flags[i]
		Cmd.Flags().BoolVarP(&f.Value, f.Name, f.ShortHand, f.DefaultValue, f.Description)
	}

	Cmd.Flags().SetInterspersed(false)
}

// executeWc executes the 'wc' command with given arguments and returns
// statistics and the length of the longest line.
func executeWc(args []string) ([][]int, int) {
	stats := [][]int{}
	longestLine := -1
	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			log.Print(err.Error())
		}
		defer file.Close()

		fileStats := []int{}
		for _, f := range flags {
			if !f.Value || f.Handler == nil {
				continue
			}

			if f.Name == "longest" {
				longestLine = max(longestLine, f.Handler(file))
				fileStats = append(fileStats, longestLine)
			} else {
				fileStats = append(fileStats, f.Handler(file))
			}
		}

		stats = append(stats, fileStats)
	}

	return stats, longestLine
}

// printStats prints statistics based on given args and stats.
func printStats(args []string, stats [][]int, longestLine int) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	printHeaders(w)

	if len(args) == 1 {
		printValues(w, stats[0], args[0])

		w.Flush()
		return
	}

	for i := range stats {
		printValues(w, stats[i], args[i])
	}

	total := calculateTotal(stats, longestLine)
	printValues(w, total, "total")

	w.Flush()
}

// calculateTotal calculates the total of each column in stats.
func calculateTotal(stats [][]int, longestLine int) []int {
	n := len(stats[0])
	total := make([]int, n)
	for i := range stats {
		for j, s := range stats[i] {
			total[j] += s
		}
	}
	if longestLine > -1 {
		total[n-1] = longestLine
	}

	return total
}

// printHeaders prints the headers based on flags.
func printHeaders(w io.Writer) {
	headers := []string{}
	for _, f := range flags {
		if f.Value {
			headers = append(headers, f.Name)
		}
	}
	headers = append(headers, "")

	for _, h := range headers {
		fmt.Fprintf(w, "%s\t", h)
	}
	fmt.Fprintln(w)
}

// printValues prints the values followed by the filename or "total".
func printValues[T any](w io.Writer, vals []T, appendix string) {
	for _, v := range vals {
		fmt.Fprintf(w, "%v\t", v)
	}
	fmt.Fprintf(w, "%s\t", appendix)
	fmt.Fprintln(w)
}
