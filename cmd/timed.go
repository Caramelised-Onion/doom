package cmd

import (
	"caramelised-onion/doom/internal/guessing"
	"context"
	"flag"
	"os"
	"sync/atomic"
	"time"
)

func RunTimed() {
	timedCmd := flag.NewFlagSet("timed", flag.ExitOnError)
	timeLimit := timedCmd.Duration("limit", 1*time.Minute, "answer as many questions right as possible in this amount of time")
	timedCmd.Parse(os.Args[2:])

	logger.Printf("Your time of %s starts now", timeLimit)
	ctx, cancel := context.WithTimeout(context.Background(), *timeLimit)
	defer cancel()

	var numCorrect atomic.Int32
	var numTotal atomic.Int32

	go func() {
		for {
			res := guessing.Guess()

			if res == guessing.Correct {
				logger.Println("Correct")
				numCorrect.Add(1)
			} else {
				logger.Println("Wrong")
			}

			numTotal.Add(1)
		}
	}()

	<-ctx.Done()
	logger.Printf("times up")
	logger.Printf("You got %d correct and %d wrong", numCorrect.Load(), numTotal.Load()-numCorrect.Load())
}
