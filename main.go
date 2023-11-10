package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func getBirthdays() (map[int]int, map[string][]string) {
	file, err := os.Open("birthdays.csv")
	if err != nil {
		fmt.Println("File could not be opened", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var birthdays []string

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			birthdays = append(birthdays, line)
			break
		} else if err != nil {
			fmt.Println("Error reading the file", err)
			break
		}

		birthdays = append(birthdays, line)
	}

	bdayCount := make(map[int]int)

	monthWithBirthdays := map[string][]string{
		"January":   {},
		"February":  {},
		"March":     {},
		"April":     {},
		"May":       {},
		"June":      {},
		"July":      {},
		"August":    {},
		"September": {},
		"October":   {},
		"November":  {},
		"December":  {},
	}

	// fmt.Println(monthWithBirthdays)

	for _, val := range birthdays {

		month := strings.Split(val, "/")

		monthInt, err := strconv.Atoi(month[0])
		if err != nil {
			continue
		}

		if value, ok := bdayCount[monthInt]; ok {
			bdayCount[monthInt] = value + 1
		} else {
			bdayCount[monthInt] = 1
		}

		monthStr := time.Month(monthInt)

		// fmt.Println(monthWithBirthdays[monthStr.String()])
		if slice, ok := monthWithBirthdays[monthStr.String()]; ok {
			monthWithBirthdays[monthStr.String()] = append(slice, val)
		}

	}

	return bdayCount, monthWithBirthdays

}

func main() {

	birthdays, monthsWithBirthdays := getBirthdays()

	fmt.Println(monthsWithBirthdays)

	commands := map[string]string{
		"[help]": "- Display a list of commands",
		"[most]": "- Display the month with the most birthdays",
		"[exit]": "- End the program",
	}

	var userInput string
	loop := true

	fmt.Println("Welcome to the birthday metrics CLI tool")
	fmt.Println("For a list of commands please type [help]")

	for loop {
		fmt.Scan(&userInput)
		switch userInput {
		case "help":
			for key, value := range commands {
				fmt.Println(key, value)
			}
		case "most":
			most := mostBirthdays(birthdays)
			fmt.Println("The month with the most birthdays is:", most)
		case "exit":
			fmt.Println("Goodbye!")
			loop = false
		default:
			fmt.Println("This is not a valid command")
		}
	}
}

func mostBirthdays(birthdays map[int]int) string {
	var keyPlaceholder, largest int
	for key, value := range birthdays {
		if value > largest {
			keyPlaceholder = key
			largest = value
		}
	}

	monthString := time.Month(keyPlaceholder)

	return monthString.String()
}
