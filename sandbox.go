// Hello project main.go
package main

import (
	"fmt"
)

func main() {
	/*s  := "Hello world"

	var freqMap = make(map[string]int)
	i := 0
	for i < len(s) {
		letter := string(s[i])
		if letter != " " {
			freqMap[letter] = freqMap[letter] + 1
		}
		i++
	}

	fmt.Println(freqMap)*/

	var point = new(mystruct)
	point.x = 11
	point.y = 12
	fmt.Print(point)
	fmt.Print(*point)
}

type mystruct struct {
	x int
	y int
}
