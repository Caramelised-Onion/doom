package main

import (
	"caramelised-onion/doom/cmd"
	"log"
	"os"
)

/*
doom stats
		- last X average
		- global average
doom practice
doom game no.
doom game gamemode parameter
doom set gamemodes no. guesses, date ranges



*/

var logger = log.New(os.Stdout, "", 0)

func main() {
	if len(os.Args) < 2 {
		logger.Fatal("expected a subcommand")
	}

	switch os.Args[1] {
	case "stats":
		logger.Println("Showing stats")
	case "practice":
		cmd.RunPractice()
	case "timed":
		cmd.RunTimed()
	default:
		logger.Fatalf("Unrecognised subcommand %s", os.Args[1])
	}
}
