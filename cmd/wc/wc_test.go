package wc

import (
	"log"
	"testing"

	"github.com/skraio/unix-utilities/internal/assert"
	"github.com/skraio/unix-utilities/internal/testutils"
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
			dummyFileName, cleanup := testutils.CreateDummyFile(t, tt.text)
			defer cleanup()

			file, err := testutils.OpenDummyFile(t, dummyFileName)
			if err != nil {
				log.Print(err.Error())
				return
			}
			defer file.Close()

			ans, err := lineCounter(file)
			if err != nil {
				log.Print(err.Error())
				return
			}

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
			dummyFileName, cleanup := testutils.CreateDummyFile(t, tt.text)
			defer cleanup()

			file, err := testutils.OpenDummyFile(t, dummyFileName)
			if err != nil {
				log.Print(err.Error())
				return
			}
			defer file.Close()

			ans, err := wordCounter(file)
			if err != nil {
				log.Print(err.Error())
				return
			}

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
			dummyFileName, cleanup := testutils.CreateDummyFile(t, tt.text)
			defer cleanup()

			file, err := testutils.OpenDummyFile(t, dummyFileName)
			if err != nil {
				log.Print(err.Error())
				return
			}
			defer file.Close()

			ans, err := byteCounter(file)
			if err != nil {
				log.Print(err.Error())
				return
			}

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
			dummyFileName, cleanup := testutils.CreateDummyFile(t, tt.text)
			defer cleanup()

			file, err := testutils.OpenDummyFile(t, dummyFileName)
			if err != nil {
				log.Print(err.Error())
				return
			}
			defer file.Close()

			ans, err := longestLine(file)
			if err != nil {
				log.Print(err.Error())
				return
			}

			assert.Equal(t, ans, tt.want)
		})
	}
}
