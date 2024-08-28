package tetrino

import (
	"bufio"
	"fmt"
	"os"
)

// Reads the file and returns all tetrinoes that were entered within the file.
func Read(s string) ([][]string, error) {
	file, err := os.Open(s)
	if err != nil {
		return nil, fmt.Errorf("ERROR")
	}
	scanner := bufio.NewScanner(file)
	t := [][]string{}
	tl := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 4 {
			tl = append(tl, line)
		} else if line != "" && len(line) != 4 {
			return nil, fmt.Errorf("ERROR")
		}
		if line == "" {
			t = append(t, tl)
			tl = []string{}
		} else if line != "" && len(tl) > 4 {
			return nil, fmt.Errorf("ERROR")
		}
	}
	if len(tl) > 0 {
		t = append(t, tl)
	}
	file.Close()
	return t, nil
}
