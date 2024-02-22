package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func preprocessOptions(options string) string {
	validOpts := map[rune]bool{'c': true, 'l': true, 'm': true, 'w': true, 'L': true}
	for _, opt := range options {
        if _, ok := validOpts[opt]; !ok {
			fmt.Printf("Invalid option: '%s'\n", string(opt))
			os.Exit(1)
		}
	}

    occurOpts := map[rune]bool{'c': false, 'l': false, 'm': false, 'w': false, 'L': false}
    for _, opt := range options {
        if !occurOpts[opt] {
            occurOpts[opt] = true
        }
    }

    optsOrder := "lwcmL"
    var processedOpts string
    for _, ch := range optsOrder {
        if occurOpts[ch] {
            processedOpts += string(ch)
        }
    }
    return processedOpts
}

func lineCounter(f *os.File) int {
	_, err := f.Seek(0, 0)
	if err != nil {
		log.Print(err.Error())
	}

	count := 0
	newLineChar := []byte{'\n'}
	for {
		buf := make([]byte, bufio.MaxScanTokenSize)
		r, err := f.Read(buf)
		if err == io.EOF {
			break
		}

		count += bytes.Count(buf[:r], newLineChar)
	}

	return count
}

func wordCounter(f *os.File) int {
	_, err := f.Seek(0, 0)
	if err != nil {
		log.Print(err.Error())
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanWords)

	count := 0
	for fileScanner.Scan() {
		count++
	}

	return count
}

func charCounter(f *os.File) int {
	_, err := f.Seek(0, 0)
	if err != nil {
		log.Print(err.Error())
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanRunes)

	count := 0
	for fileScanner.Scan() {
		count++
	}

	return count
}

func byteCounter(f *os.File) int {
	_, err := f.Seek(0, 0)
	if err != nil {
		log.Print(err.Error())
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanBytes)

	count := 0
	for fileScanner.Scan() {
		count++
	}

	return count
}

func longestLine(f *os.File) int {
    _, err := f.Seek(0, 0)
    if err != nil {
        log.Print(err.Error())
    }

    fileScanner := bufio.NewScanner(f)
    fileScanner.Split(bufio.ScanLines)

    maxLen := 0
    for fileScanner.Scan() {
        maxLen = max(maxLen, len(fileScanner.Text()))
    }

    return maxLen
}
