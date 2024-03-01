package wc

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// lineCounter counts the number of lines in a file.
func lineCounter(f *os.File) (int, error) {
	_, err := f.Seek(0, 0)
	if err != nil {
        return 0, err
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

	return count, nil
}

// wordCounter counts the number of words in a file.
func wordCounter(f *os.File) (int, error) {
	_, err := f.Seek(0, 0)
	if err != nil {
        return 0, err
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanWords)

	count := 0
	for fileScanner.Scan() {
		count++
	}

	return count, nil
}

// byteCounter counts the number of bytes in a file.
func byteCounter(f *os.File) (int, error) {
	_, err := f.Seek(0, 0)
	if err != nil {
        return 0, err
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanBytes)

	count := 0
	for fileScanner.Scan() {
		count++
	}

	return count, nil
}

// longestLine finds the length of the longest line in a file.
func longestLine(f *os.File) (int, error) {
	_, err := f.Seek(0, 0)
	if err != nil {
        return 0, err
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	maxLen := 0
	for fileScanner.Scan() {
		maxLen = max(maxLen, len(fileScanner.Text()))
	}

	return maxLen, nil
}
