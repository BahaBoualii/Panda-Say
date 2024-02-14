package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes!")
		fmt.Println("Usage: forune | pandasay")
		return
	}

	var lines []string

	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}

	var panda = `
	\
	 \
	  \
	   \
	    \
	_§§§§§___§§§§§§§__§§§§§§
	§§§§§§§§§_______§§§§§§c§§
	§§§§§§_____________§§§§d§
	§§§§_________________§§§§
	§§§___________________§§§
	_§§___________________§§
	_§_____________________§
	_§__§§§___________§§§__§
	§__§___§_________§___§__§
	§_§_§§__§_______§__§§_§_§
	§_§_§§§_§_______§_§§§_§_§
	§_§_§§§_§_§§§§§_§_§§§_§_§
	§__§___§__§§§§§__§___§__§
	_§__§§§____§§§____§§§__§
	_§_______§__§__§_______§
	__§_______§§_§§_______§
	___§§_______________§§
	_____§§___________§§
	____§§§§§§§§§§§§§§§§§
	___§§_______________§§
	__§§______§§§§§______§§
	_§§§______§___§______§§§
	§§§§§_____§___§_____§§§§§
	§§§§§____§_____§____§§§§§
	§§§§§____§_____§____§§§§§
	_§§§§§§§§_______§§§§§§§§
	__§§§§§§§_______§§§§§§§
	__§§___§§_______§§___§§
	__§§___§§_______§§___§§
	__§§___§§_§§§§§_§§___§§
	__§§§§§§§§_____§§§§§§§§
	___§§§§§_________§§§§§
	
	`

	lines = TabsToSpaces(lines)
	maxwidth := CalculateMaxWidth(lines)
	messages := NormalizeStringLength(lines, maxwidth)
	balloon := BuildBalloon(messages, maxwidth)
	fmt.Println(balloon)
	fmt.Println(panda)
	fmt.Println()
}
