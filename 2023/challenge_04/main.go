package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func countReps(text string, char string) int {
	count := 0
	for i := 0; i < len(text); i++ {
		if text[i] == []byte(char)[0] {
			count++
		}
	}
	return count
}

func main() {
	valids := 0
	re := regexp.MustCompile(`(\w+)-(\w+)`)
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		this := scanner.Text()
		text := re.FindStringSubmatch(this)
		if text != nil {
			checksum := ""
			for _, char := range text[1] {
				if countReps(text[1], string(char)) == 1 {
					checksum += string(char)
				}
			}
			if checksum == text[2] {
				valids++
				if valids == 33 {
					fmt.Println("Solución al reto 4:\nsubmit " + text[2])
					break
				}
			}
		}
	}
}
