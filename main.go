package main

import (
	"fmt"
)

func main() {
	fmt.Println("Шахматная доска")
	fmt.Println("Введите ширину доски:")
	var width int
	fmt.Scan(&width)

	fmt.Println("Введите высоту доски:")
	var height int
	fmt.Scan(&height)
	var isFirstBlack bool
	var line string
	for j := 0; j < height; j++ {
    line = ""
		if j%2 == 0 {
			isFirstBlack = true
		} else {
			isFirstBlack = false
		}
		for i := 0; i < width; i++ {
			if isFirstBlack == (i%2 == 0) {
				line += " "
			} else {
				line += "*"
			}
			
		}
    fmt.Println(line)
	}
}
