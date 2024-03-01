package ls

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/skraio/unix-utilities/internal/assert"
	"github.com/skraio/unix-utilities/internal/testutils"
)

func TestLongFormat(t *testing.T) {
	tests := []struct {
		name string
		text []byte
		want FileAttributes
	}{
		{
			name: "Short file",
			text: []byte("Without just one nest\nA bird can call the world home\nLife is your career\n"),
			want: FileAttributes{
				fileMode: os.FileMode(0644),
				ulink:    "1",
				uid:      "tdk",
				fileSize: "73",
				modTime:  time.Now().Format("Jan _2 15:04"),
			},
		},
		{
			name: "Empty file",
			text: []byte(""),
			want: FileAttributes{
				fileMode: os.FileMode(0644),
				ulink:    "1",
				uid:      "tdk",
				fileSize: "0",
				modTime:  time.Now().Format("Jan _2 15:04"),
			},
		},
	}

	for _, tt := range tests {
		dummyFileName, cleanup := testutils.CreateDummyFile(t, tt.text)
		defer cleanup()

		file, err := testutils.OpenDummyFile(t, dummyFileName)
		if err != nil {
			log.Print(err.Error())
			return
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			t.Fatal(err)
		}

		ans, err := longFormat(fileInfo)
		if err != nil {
			log.Print(err.Error())
			return
		}

		assert.Equal(t, ans, tt.want)
		// cleanup()
	}
}

func TestHumanReadableSize(t *testing.T) {
	tests := []struct {
		name string
		size int64
		want string
	}{
		{
			name: "Small size",
			size: 4096,
			want: "4096",
		},
		{
			name: "Big size",
			size: 5*1024*1024 + 12345,
			want: "5M",
		},
	}

	for _, tt := range tests {
		ans := humanReadableSize(tt.size)

		assert.Equal(t, ans, tt.want)
	}
}
