// Package cat provides functionality for concatenating files and printing
// their content.
package cat

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/skraio/unix-utilities/cmdflags"
	"github.com/spf13/cobra"
)

// content represents the content to be printed by cat.
type content struct {
	lineNumber []string
	text       []string
}

// catFlags represents the flags used by the cat command.
type catFlags struct {
	squeezeBlank   bool
	endOfLine      bool
	numberNonblank bool
	number         bool
}

var pFlags catFlags

// flags definition for cat command.
var flags = []cmdflags.Flag{
	{Value: &pFlags.squeezeBlank, Name: "squeeze-blank", ShortHand: "s", DefaultValue: false, Description: "squeeze multiple adjacent blank lines"},
	{Value: &pFlags.endOfLine, Name: "end-line-chars", ShortHand: "e", DefaultValue: false, Description: "display end-of-line characters $"},
	{Value: &pFlags.numberNonblank, Name: "number-nonblank", ShortHand: "b", DefaultValue: false, Description: "number non-blank output lines"},
	{Value: &pFlags.number, Name: "number", ShortHand: "n", DefaultValue: false, Description: "number all output lines"},
}

// Cmd represents the 'cat' command configuration using Cobra.
var Cmd = &cobra.Command{
	Use:   "cat [-f flags] [file]...",
	Short: "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cont := &content{}
		err := cont.executeCat(args)
		if err != nil {
			log.Print(err.Error())
			return
		}
	},
}

// init initializes the 'cat' command by setting up flags.
func init() {
	cmdflags.ParseFlags(flags, Cmd)
}

// executeLs executes the cat command with given arguments.
func (cont *content) executeCat(args []string) error {
	startIdx := 0
	for _, arg := range args {
		err := cont.execute(arg, startIdx)
		if err != nil {
			return err
		}
	}

	if pFlags.numberNonblank || pFlags.number {
		cont.numberLines()
	}

	cont.printText()
	return nil
}

// execute reads the content of the file and stores it in the content struct.
func (cont *content) execute(filename string, startIdx int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fileContent, err := readFileContent(file)
	if err != nil {
		return err
	}

	if pFlags.squeezeBlank {
		cont.text = append(cont.text, squeezeBlankLines(fileContent)...)
	} else {
		cont.text = append(cont.text, fileContent...)
	}

	return nil
}

// readFileContent reads the content of a file.
func readFileContent(f *os.File) ([]string, error) {
	_, err := f.Seek(0, 0)
	if err != nil {
		return []string{}, err
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	text := []string{}
	for fileScanner.Scan() {
		text = append(text, fileScanner.Text())
	}

	return text, nil
}

// numberLines numbers the lines based on the flags.
func (cont *content) numberLines() {
	if pFlags.numberNonblank {
		cont.numberNonblankLines()
	} else if pFlags.number {
		cont.numberAllLines()
	}
}

// printText prints the content stored in the content struct.
func (cont *content) printText() {
	if len(cont.lineNumber) == 0 {
		for _, l := range cont.text {
			fmt.Println(l)
		}
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight)
	defer w.Flush()

	for i := range cont.lineNumber {
		fmt.Fprintf(w, "%s\t\t%s", cont.lineNumber[i], cont.text[i])
		fmt.Fprintln(w)
	}
}
