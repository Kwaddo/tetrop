package main

import (
	"fmt"
	"log"
	"math"
	"os"
	fn "tetrino/pkg"
)

// The one and only.
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please use the proper inputs!\n[go run . <TEXTFILE>]")
	}
	t, err := fn.Read(os.Args[1])
	if err != nil {
		log.Fatal("File does not exist!")
	}
	if !fn.ValidCheck(t) || len(t) == 0 {
		log.Fatal("Invalid Tetominos!")
	}
	bs := int(math.Ceil(math.Sqrt(float64(len(t) * 4))))
	b := make([][]string, bs)
	for i := range b {
		b[i] = make([]string, bs)
		for j := range b[i] {
			b[i][j] = "."
		}
	}
	ct := fn.CutLines(t)
	end := fn.Solve(ct, b)
	for _, row := range end {
		var s string
		for _, char := range row {
			s += char
		}
		fmt.Println(s)
	}
}
