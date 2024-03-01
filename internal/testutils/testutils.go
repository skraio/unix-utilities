package testutils

import (
	"log"
	"os"
	"testing"
)

func CreateDummyFile(t *testing.T, test []byte) (string, func()) {
	dummyFileName := "dummy_file.txt"
	err := os.WriteFile(dummyFileName, test, 0644)
	if err != nil {
		log.Fatalf("Error creating dummy file: %v", err)
	}

	return dummyFileName, func() {
		err := os.Remove(dummyFileName)
		if err != nil {
			// t.Fatalf("Error removing dummy file: %v", err)
            return
		}
	}
}

func OpenDummyFile(t *testing.T, dummyFileName string) (*os.File, error) {
	file, err := os.Open(dummyFileName)
	if err != nil {
        return nil, err
	}

	return file, nil
}
