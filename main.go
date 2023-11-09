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

func getBirthdays() map[int]int {
	// Open the file we want to read
	file, err := os.Open("birthdays.csv")
	if err != nil {
		fmt.Println("File could not be opened", err)
	}
	defer file.Close()

	// create a new buffered reader
	reader := bufio.NewReader(file)

	// Creating a slice and add all the birthdays to be looped through
	var birthdays []string

	// Infinite loop that loops through the file looking for new lines and prints the lines unless it is the EOF (end of file) or we get an err
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

	// Looping through the created birthdays slice
	for _, val := range birthdays {

		// Seperates the month from the rest of the date
		month := strings.Split(val, "/")
		// fmt.Println(month[0])

		// Converts the month string into an integer using the strconv package
		monthInt, err := strconv.Atoi(month[0])
		if err != nil {
			// fmt.Println("There was a  conversiong problem", err)
			continue
		}

		// Check to see if the month exists in the csv file and if it does add one to its value
		if value, ok := bdayCount[monthInt]; ok {
			bdayCount[monthInt] = value + 1
		} else {
			bdayCount[monthInt] = 1
		}
	}

	return bdayCount

}

func main() {

	birthdays := getBirthdays()

	commands := map[string]string{
		"[help]": "- display a list of commands",
		"[most]": "- display the month with the most birthdays",
		"[exit]": "- End the program",
	}

	var userInput string
	loop := true

	// This is the cli loop where the programs askes for user input
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

// takes the map of birhtdays and loops through to find the month with the most birthdays via its value
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
