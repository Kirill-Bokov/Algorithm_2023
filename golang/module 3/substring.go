package module3

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func Substring() {
	input, _ := os.OpenFile("input.txt", os.O_RDONLY, 0666)
	defer input.Close()
	reader := bufio.NewReader(input)
	var line []string
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		str = strings.TrimSuffix(str, "\r")
		str = strings.TrimSuffix(str, "\n")
		line = append(line, str)
	}
	basis := strings.TrimSuffix(strings.TrimSuffix(line[0], "\n"), "\r")
	pattern := strings.TrimSuffix(strings.TrimSuffix(line[1], "\n"), "\r")
	output, _ := os.OpenFile("output.txt", os.O_APPEND|os.O_TRUNC|os.O_CREATE, 0644)
	output.WriteString(foundindexes(basis, pattern))
}

func foundindexes(basis, pattern string) string {
	var output string
	for i := 0; i < len(basis)-len(pattern); i++ {
		found := true
		for j := 0; j < len(pattern); j++ {
			if basis[i+j] != pattern[j] {
				found = false
				break
			}
		}
		if found {
			if len(output) > 0 {
				output += " " + strconv.Itoa(i)
			} else {
				output += strconv.Itoa(i)
			}

		}

		break
	}
	return output
}
