package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isIdValid(id string) bool {
	if id == "" {
		return false
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if re.MatchString(id) {
		return true
	}
	return false
}

func isEmailValid(email string) bool {
	if email == "" {
		return false
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+\.[a-zA-z]{2,3}$`)
	if re.MatchString(email) {
		return true
	}
	return false
}

func isUserValid(user string) bool {
	if user == "" {
		return false
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if re.MatchString(user) {
		return true
	}
	return false
}

func isLocationValid(location string) bool {
	if location == "" {
		return true
	} else {
		re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
		if re.MatchString(location) {
			return true
		}
	}
	return false
}

func isAgeValid(age string) bool {
	if age == "" {
		return true
	} else {
		_, err := strconv.Atoi(age)
		return err == nil
	}
	return false
}

func main() {
	file, err := os.Open("./input.txt")
	message := ""
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), ",")
		if !isIdValid(items[0]) || !isUserValid(items[1]) ||
			!isEmailValid(items[2]) || !isAgeValid(items[3]) ||
			!isLocationValid(items[4]) {
			message += string(items[1][0])
		}
	}
	fmt.Println("Solución al reto 5:\nsubmit " + message)
}
