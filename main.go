package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var inputTime time.Time

	timeString, err := getInputTime()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Parse input string to time.Time with time.Kitchen layout
	inputTime, err = time.Parse(time.Kitchen, timeString)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Get local time zone
	zone, _ := time.Now().Zone()

	fmt.Println("Local Time:", inputTime.Local().Format(time.Kitchen), zone)
}

// Function to read user input from the Command-line
func getInputTime() (string, error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("*****************************")
	fmt.Println("UTC to Local Time Converter")
	fmt.Println("*****************************")
	fmt.Println("Enter the time in UTC below - (Format HH:MM PM/AM)")
	fmt.Println("-----------------------------")

	for {
		fmt.Print("-> ")
		time, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		} else {
			// convert CRLF to LF
			time = strings.Replace(time, "\r\n", "", -1)

			// trim whitespaces
			time = strings.Replace(time, " ", "", -1)

			return time, nil
		}
	}
}
