package guessing

import (
	"caramelised-onion/doom/internal/date"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var logger = log.New(os.Stdout, "", 0)

type GuessResult int

const (
	Correct GuessResult = iota
	Wrong
	Quit
)

func Guess() GuessResult {
	actualDate := date.RandDate()
	logger.Printf("What day of the week was %s?", actualDate.Format("2006-01-02"))

	var userInput string
	fmt.Scanln(&userInput)

	if userInput == "q" {
		return Quit
	}

	dayAsInt, err := strconv.Atoi(userInput)
	if err != nil {
		logger.Printf("Input %s could not be parsed as a valid day of the week", userInput)
		return Wrong
	}

	guess := time.Weekday(dayAsInt)

	if actualDate.Weekday() == guess {
		return Correct
	} else {
		return Wrong
	}
}
