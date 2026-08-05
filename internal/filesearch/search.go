package filesearch

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Search(filepath string, search string) ([]string, error) {

	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer func() { _ = file.Close() }()

	reader := bufio.NewReader(file)
	endOfLine := '\n'

	occurrences := make([]string, 0)
	var lineNumber int = 1

	for {
		buf, err := reader.ReadBytes(byte(endOfLine))

		if len(buf) > 0 {
			sentence := string(buf)
			if strings.Contains(sentence, search) {
				occurrences = append(occurrences, formatLine(lineNumber, buf))
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		lineNumber++
	}

	return occurrences, nil
}

func formatLine(lineNumber int, line []byte) string {
	return fmt.Sprintf("line %d: %s", lineNumber, string(line))
}
