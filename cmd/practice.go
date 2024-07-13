package cmd

import (
	"caramelised-onion/doom/internal/guessing"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", 0)

func RunPractice() {
	currentRound := 1
	stillPlaying := true
	numCorrect := 0

gameLoop:
	for stillPlaying {
		res := guessing.Guess()

		switch res {
		case guessing.Quit:
			break gameLoop
		case guessing.Correct:
			numCorrect++
			logger.Println("Correct")
		case guessing.Wrong:
			logger.Println("Wrong!")
		}

		currentRound++
	}

	logger.Printf("You got %d correct and %d wrong", numCorrect, currentRound-numCorrect)
}
