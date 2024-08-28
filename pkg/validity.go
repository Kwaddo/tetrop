package tetrino

// This function focuses on cutting any lines that aren't required and are helpful to slim the board.
func CutLines(t [][]string) [][]string {
	ct := [][]string{}
	for _, tetromino := range t {
		var nt []string
		for _, line := range tetromino {
			count := 0
			for _, char := range line {
				if char == '#' {
					count++
				}
			}
			if count > 0 {
				nt = append(nt, line)
			}
		}
		ct = append(ct, nt)
	}
	for i := range ct {
		for j := len(ct[i][0]) - 1; j >= 0; j-- {
			count := 0
			for k := range ct[i] {
				if ct[i][k][j] == '#' {
					count++
				}
			}
			if count == 0 {
				for l := range ct[i] {
					ct[i][l] = Remove(ct[i][l], j)
				}
			}
		}
	}
	return ct
}

// A function to remove the character found at the given index.
func Remove(s string, index int) string {
	if index < 0 || index >= len(s) {
		return s
	}
	runes := []rune(s)
	runes = append(runes[:index], runes[index+1:]...)
	return string(runes)
}

// This function checks of the given tetrino is valid.
func ValidCheck(t [][]string) bool {
	for _, tr := range t {
		connections := 0
		count := 0
		for iv, line := range tr {
			for ih, char := range line {
				hc := 0
				if char != '#' && char != '.' {
					return false
				} else if char == '#' {
					count++
					if iv > 0 && tr[iv-1][ih] == '#' {
						hc++
					}
					if iv < len(tr)-1 && tr[iv+1][ih] == '#' {
						hc++
					}
					if ih > 0 && tr[iv][ih-1] == '#' {
						hc++
					}
					if ih < len(line)-1 && tr[iv][ih+1] == '#' {
						hc++
					}
					if hc == 0 {
						return false
					} else {
						connections += hc
					}
				}
			}
		}
		if connections < 6 || count > 4 {
			return false
		}
	}
	return true
}
