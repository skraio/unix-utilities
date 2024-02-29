package wc

import (
	"log"
	"os"
	"testing"

	"github.com/skraio/unix-utilities/internal/assert"
)

func TestLineCounter(t *testing.T) {
	tests := []struct {
		name string
		text []byte
		want int
	}{
		{
			name: "Short file",
			text: []byte("Without just one nest\nA bird can call the world home\nLife is your career\n"),
			want: 3,
		},
		{
			name: "Empty file",
			text: []byte(""),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dummyFileName, cleanup := createDummyFile(t, tt.text)
			defer cleanup()

			file := openDummyFile(t, dummyFileName)
			defer file.Close()

			ans := lineCounter(file)

			assert.Equal(t, ans, tt.want)
		})
	}
}

func TestWordCounter(t *testing.T) {
	tests := []struct {
		name string
		text []byte
		want int
	}{
		{
			name: "Short file",
			text: []byte("Without just one nest\n\nA bird can call the world home\n\nLife is your career\n"),
			want: 15,
		},
		{
			name: "Empty file",
			text: []byte(""),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dummyFileName, cleanup := createDummyFile(t, tt.text)
			defer cleanup()

			file := openDummyFile(t, dummyFileName)
			defer file.Close()

			ans := wordCounter(file)

			assert.Equal(t, ans, tt.want)
		})
	}
}

func TestByteCounter(t *testing.T) {
	tests := []struct {
		name string
		text []byte
		want int
	}{
		{
			name: "Short file",
			text: []byte("Hello\n\nworld\n\n"),
			want: 14,
		},
		{
			name: "Empty file",
			text: []byte(""),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dummyFileName, cleanup := createDummyFile(t, tt.text)
			defer cleanup()

			file := openDummyFile(t, dummyFileName)
			defer file.Close()

			ans := byteCounter(file)

			assert.Equal(t, ans, tt.want)
		})
	}
}

func TestLongestLine(t *testing.T) {
	tests := []struct {
		name string
		text []byte
		want int
	}{
		{
			name: "Short file",
			text: []byte("Without just one nest\nA bird can callthe world homeLife is your\n career\n"),
			want: 41,
		},
		{
			name: "Empty file",
			text: []byte(""),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            dummyFileName, cleanup := createDummyFile(t, tt.text)
            defer cleanup()

            file := openDummyFile(t, dummyFileName)
            defer file.Close()

			ans := longestLine(file)

			assert.Equal(t, ans, tt.want)
		})
	}
}

func createDummyFile(t *testing.T, test []byte) (string, func()) {
	dummyFileName := "dummy_file.txt"
	err := os.WriteFile(dummyFileName, test, 0644)
	if err != nil {
		log.Fatalf("Error creating dummy file: %v", err)
	}

	return dummyFileName, func() {
		err := os.Remove(dummyFileName)
		if err != nil {
			t.Fatalf("Error removing dummy file: %v", err)
		}
	}
}

func openDummyFile(t *testing.T, dummyFileName string) *os.File {
	file, err := os.Open(dummyFileName)
	if err != nil {
		log.Print(err.Error())
	}

	return file
}
