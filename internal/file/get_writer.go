package file

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func GetWriter(defaultWriter io.Writer, outputFile string) (io.Writer, func() error, error) {
	nop := func() error { return nil }
	path := strings.TrimSpace(outputFile)

	switch len(path) {
	case 0:
		return defaultWriter, nop, nil

	default:
		outputFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

		if err != nil {
			return nil, nop, fmt.Errorf("unable to create report file: %w", err)
		}

		return outputFile, func() error {
			return outputFile.Close()
		}, nil
	}
}
