package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	invalids := 0
	re := regexp.MustCompile(`(\d+)-(\d+) (\w): (.+)`)
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
			min, _ := strconv.Atoi(text[1])
			max, _ := strconv.Atoi(text[2])
			char := text[3]
			key := text[4]

			reps := countReps(key, char)
			if reps < min || reps > max {
				invalids++
				if invalids == 42 {
					fmt.Println("Solución al reto 2:\nsubmit " + key)
					break
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
