package cat

import "strconv"

// squeezeBlankLines squeezes the multiple adjacent blank lines from the given
// text slice.
func squeezeBlankLines(text []string) []string {
	result := []string{}
	result = append(result, text[0])

	for i := 1; i < len(text); i++ {
		if text[i] == "" && text[i-1] == "" {
			continue
		}
		result = append(result, text[i])
	}

	return result
}

// numberNonblankLines numbers non-blank lines in the content structure.
func (cont *content) numberNonblankLines() {
	lineCounter := 1

	for _, line := range cont.text {
		if line != "" {
			cont.lineNumber = append(cont.lineNumber, strconv.Itoa(lineCounter))
			lineCounter++
		} else {
			cont.lineNumber = append(cont.lineNumber, "")
		}
	}
}

// numberAllLines nubmers all lines in the content structure.
func (cont *content) numberAllLines() {
	cont.lineNumber = make([]string, len(cont.text))

	for i := range cont.text {
		cont.lineNumber[i] = strconv.Itoa(i + 1)
	}
}
