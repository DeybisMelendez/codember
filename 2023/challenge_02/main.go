package main

import (
	"fmt"
	"strconv"
	"strings"
)

const INPUT string = "&###@&*&###@@##@##&######@@#####@#@#@#@##@@@@@@@@@@@@@@@*&&@@@@@@@@@####@@@@@@@@@#########&#&##@@##@@##@@##@@##@@##@@##@@##@@##@@##@@##@@##@@##@@##@@##@@&"

func main() {
	var value int = 0
	var solution string = ""
	for _, char := range strings.Split(INPUT, "") {
		switch char {
		case "#":
			value++
		case "@":
			value--
		case "*":
			value = value * value
		case "&":
			solution += strconv.Itoa(value)
		}
	}
	fmt.Println("Solución al reto 2:\nsubmit " + solution)
}
