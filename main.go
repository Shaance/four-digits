package main

import (
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func getUserInput(msg string) []int {
	var userInput string
	err := validateInput(userInput)

	for err != nil {
		// first iteration will not trigger error message
		if userInput != "" {
			fmt.Println(err)
		}

		fmt.Println(msg)
		fmt.Scanln(&userInput)
		err = validateInput(userInput)
	}

	intArrayOutput := make([]int, 4)

	for i := range userInput {
		intArrayOutput[i], _ = strconv.Atoi(string(userInput[i]))
	}

	return intArrayOutput
}

func getInputComparison(userInput []int, answer []int) string {
	a, b := 0, 0
	answerNumbers := make(map[int]bool)
	for _, ansNb := range answer {
		answerNumbers[ansNb] = true
	}

	for i, userNb := range userInput {
		if userNb == answer[i] {
			a += 1
		} else if _, exists := answerNumbers[userNb]; exists {
			b += 1
		}
	}

	return fmt.Sprintf("%dA%dB", a, b)
}

func isNumeric(s string) bool {
	return regexp.MustCompile(`^[0-9]*$`).MatchString(s)
}

func validateInput(userInput string) error {
	if len(userInput) != 4 {
		return fmt.Errorf("[Error] Input needs to be 4 characters long")
	}

	if !isNumeric(userInput) {
		return fmt.Errorf("[Error] Input needs to be composed of 4 digits between 0 and 9")
	}

	seen := make(map[rune]bool)
	for _, nb := range userInput {
		if _, exists := seen[nb]; exists {
			return fmt.Errorf("[Error] Input can't have duplicates")
		}
		seen[nb] = true
	}

	return nil
}

func gameFinished(inputComparison string) bool {
	return inputComparison == "4A0B"
}

func generateAnswer() []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(10)[:4]
}

func getTriesString(tries int) string {
	if tries == 1 {
		return "try"
	}
	return "tries"
}

func getNbTries() int {
	const defaultTries = 6
	var tries int
	flag.IntVar(&tries, "t", defaultTries, "Specify number of tries. Default is 6")
	flag.Parse()
	if tries < 1 {
		// does not make sense, just make it default
		tries = defaultTries
	}
	return tries
}

func main() {
	tries := getNbTries()
	answer := generateAnswer()
	finished := false

	for !finished && tries > 0 {
		triesString := getTriesString(tries)
		fmt.Printf("You have %d %s left.\n", tries, triesString)
		userInput := getUserInput("Input your 4 digit guess:")
		inputComparison := getInputComparison(userInput, answer)
		fmt.Println(inputComparison)
		finished = gameFinished(inputComparison)
		tries -= 1
	}

	fmt.Println()
	if finished {
		fmt.Println("You won!")
	} else {
		fmt.Printf("You've lost, answer was %v\n", answer)
	}

}
