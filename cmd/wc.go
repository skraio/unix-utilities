package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const (
	numFlags = 5

	// Flags
	lineFlag        = 'l'
	wordFlag        = 'w'
	byteFlag        = 'c'
	charFlag        = 'm'
	longestLineFlag = 'L'
)

var counterFunctions = map[rune]func(*os.File) int{
	lineFlag        : lineCounter,
	wordFlag        : wordCounter,
	byteFlag        : byteCounter,
	charFlag        : charCounter,
	longestLineFlag : longestLine,
}

var wcCmd = &cobra.Command{
	Use:  "wc [file]... [-f flags]",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        if options == "" {
            options = "lwc"
        } else {
            options = preprocessOptions(options)
        }
		executeWc(args)
	},
}

func init() {
	wcCmd.Flags().StringVar(&options, "f", "", "command's flags")
}

func executeWc(args []string) {
    numOptions := len(options)
    totalStats := make([]int, numOptions)

    longestLineVal := -1

	for _, arg := range args {
		filename := arg
		file, err := os.Open(filename)
		if err != nil {
			log.Print(err.Error())
		}
        defer file.Close()

		curStats := make([]int, numOptions)
		for i, c := range options {
            if counterFunc, ok := counterFunctions[c]; ok {
                curStats[i] = counterFunc(file)

                if c == longestLineFlag {
                    longestLineVal = max(longestLineVal, curStats[i])
                }
            }
		}
        printStats(filename, curStats)

		for i := 0; i < numOptions; i++ {
			totalStats[i] += curStats[i]
		}
        if longestLineVal >= 0 {
            totalStats[numOptions - 1] = longestLineVal
        }
	}

	if len(args) > 1 {
		printStats("total", totalStats)
	}
}

func printStats(filename string, stats []int) {
	for _, stat := range stats {
		if stat != 0 {
            fmt.Print(stat, " ")
		}
	}
	fmt.Println(filename)
}
