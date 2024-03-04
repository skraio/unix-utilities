package cat

import (
	"testing"

	"github.com/skraio/unix-utilities/internal/assert"
)

func TestSqueezeBlankLines(t *testing.T) {
	tests := []struct {
		name string
		text []string
		want []string
	}{
		{
			name: "Empty input",
			text: []string{""},
			want: []string{""},
		},
		{
			name: "Multiple ajdacent blank lines",
			text: []string{"Without just one nest", "", "", "", "", "A bird can call the world home", "", "", "", "", "Life is your career", "", "", ""},
			want: []string{"Without just one nest", "", "A bird can call the world home", "", "Life is your career", ""},
		},
		{
			name: "No adjacent blank lines",
			text: []string{"Without just one nest", "", "A bird can call the world home", "", "Life is your career", ""},
			want: []string{"Without just one nest", "", "A bird can call the world home", "", "Life is your career", ""},
		},
	}

	for _, tt := range tests {

		ans := squeezeBlankLines(tt.text)
		assert.EqualStr(t, ans, tt.want)
	}
}

func TestNumberNonblankLines(t *testing.T) {
	tests := []struct {
		name string
		text []string
		want []string
	}{
		{
			name: "Empty input",
			text: []string{""},
			want: []string{""},
		},
		{
			name: "Non-blank lines",
			text: []string{"", "Without just one next", "", "", "A bird can call the world home", ""},
			want: []string{"", "1", "", "", "2", ""},
		},
		{
			name: "No non-blank lines",
			text: []string{"Without just one next", "A bird can call the world home", "Life is your career"},
			want: []string{"1", "2", "3"},
		},
	}

	for _, tt := range tests {
		ans := &content{text: tt.text}
		ans.numberNonblankLines()

		assert.EqualStr(t, ans.lineNumber, tt.want)
	}
}

func TestNumberAllLines(t *testing.T) {
	tests := []struct {
		name string
		text []string
		want []string
	}{
		{
			name: "Empty input",
			text: []string{},
			want: []string{},
		},
		{
			name: "Non-blank lines",
			text: []string{"", "Without just one next", "", "", "A bird can call the world home", ""},
			want: []string{"1", "2", "3", "4", "5", "6"},
		},
		{
			name: "Blank lines",
			text: []string{"", "", ""},
			want: []string{"1", "2", "3"},
		},
	}

	for _, tt := range tests {
		ans := &content{text: tt.text}
		ans.numberAllLines()

		assert.EqualStr(t, ans.lineNumber, tt.want)
	}
}
