package main

import (
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

type Hintable interface {
	getHint([]int) string
}

type Checkable interface {
	validate()
}
type UserInput struct {
	originalStringInput string
	hintableInput []int // to be able to call getHint
	err error
}

func (u UserInput) getHint(answer []int) string {
	a, b := 0, 0
	answerNumbers := make(map[int]bool)
	for _, ansNb := range answer {
		answerNumbers[ansNb] = true
	}

	for i, userNb := range u.hintableInput {
		if userNb == answer[i] {
			a += 1
		} else if _, exists := answerNumbers[userNb]; exists {
			b += 1
		}
	}

	return fmt.Sprintf("%dA%dB", a, b)
}

func (u *UserInput) validate() {
	userInput := u.originalStringInput
	if len(userInput) != 4 {
		u.err = fmt.Errorf("[Error] Input needs to be 4 characters long")
		return
	}

	if !isNumeric(userInput) {
		u.err = fmt.Errorf("[Error] Input needs to be composed of 4 digits between 0 and 9")
		return
	}

	seen := make(map[rune]bool)
	for _, nb := range userInput {
		if _, exists := seen[nb]; exists {
			u.err = fmt.Errorf("[Error] Input can't have duplicates")
			return
		}
		seen[nb] = true
	}

	u.err = nil
}

func getUserInput(msg string) UserInput {
	var userInput UserInput
	userInput.originalStringInput = ""
	userInput.validate()
	for userInput.err != nil {
		// does not run on first try
		if userInput.originalStringInput != "" {
			fmt.Println(userInput.err.Error())
		}

		fmt.Println(msg)
		fmt.Scanln(&userInput.originalStringInput)
		userInput.validate()
	}

	intArrayOutput := make([]int, 4)
	ogInput := userInput.originalStringInput
	for i := range ogInput {
		intArrayOutput[i], _ = strconv.Atoi(string(ogInput[i]))
	}
	userInput.hintableInput = intArrayOutput
	return userInput
}

func isNumeric(s string) bool {
	return regexp.MustCompile(`^[0-9]*$`).MatchString(s)
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
	const defaultTries = 15
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
		ui := getUserInput("Input your 4 digit guess:")
		hint := ui.getHint(answer)
		fmt.Println(hint)
		finished = gameFinished(hint)
		tries -= 1
	}

	fmt.Println()
	if finished {
		fmt.Println("You won!")
	} else {
		fmt.Printf("You've lost, answer was %v\n", answer)
	}

}
