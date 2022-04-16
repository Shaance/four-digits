package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
	"strconv"
	"flag"
)

func get_user_input(msg string) []int {
	var user_input string
	err := validate_input(user_input)

	for err != nil {
		// first iteration will not trigger error message
		if user_input != "" {
			fmt.Println(err)
		}

		fmt.Println(msg)
		fmt.Scanln(&user_input)
		err = validate_input(user_input)
	}

	int_input := make([]int, 4)
	
	for i := range user_input {
		int_input[i], _ = strconv.Atoi(string(user_input[i]))
	}

	return int_input
}

func get_input_comparison(user_input []int, answer []int) string {
	a, b := 0, 0
	answer_numbers := make(map[int]bool)
	for _, ans_nb := range answer {
		answer_numbers[ans_nb] = true
	}

	for i, u_nb := range user_input {
		if u_nb == answer[i] {
			a += 1
		} else if _, exists := answer_numbers[u_nb]; exists == true {
			b += 1
		}
	}

	return fmt.Sprintf("%dA%dB", a, b)
}

func is_numeric(s string) bool {
	return regexp.MustCompile(`^[0-9]*$`).MatchString(s)
}

func validate_input(user_input string) error {
	if len(user_input) != 4 {
		return fmt.Errorf("[Error] Input needs to be 4 characters long")
	}

	if !is_numeric(user_input) {
		return fmt.Errorf("[Error] Input needs to be composed of 4 digits between 0 and 9")
	}

	seen := make(map[rune]bool)
	for _, nb := range user_input {
		if _, ok := seen[nb]; ok == true {
			return fmt.Errorf("[Error] Input can't have duplicates")
		}
		seen[nb] = true
	}

	return nil
}

func game_finished(input_comparison string) bool {
	return input_comparison == "4A0B"
}

func generate_answer() []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(10)[:4]
}

func get_tries_string(tries int) string {
	if tries == 1 {
		return "try"
	}
	return "tries"
}

func get_nb_tries() int {
	const default_tries = 6
	var tries int
	flag.IntVar(&tries, "t", default_tries, "Specify number of tries. Default is 6")
	flag.Parse()
	if tries < 1 {
		// does not make sense, just make it default
		tries = default_tries
	}
	return tries
}

func main() {
	tries := get_nb_tries()
	answer := generate_answer()
	finished := false
	
	for !finished && tries > 0 {
		tries_string := get_tries_string(tries)
		fmt.Printf("You have %d %s left.\n", tries, tries_string)
		user_input := get_user_input("Input your 4 digit guess:")
		input_comparison := get_input_comparison(user_input, answer)
		fmt.Println(input_comparison)
		finished = game_finished(input_comparison)
		tries -= 1
	}

	fmt.Println()
	if finished {
		fmt.Println("You won!")
	} else {
		fmt.Printf("You've lost, answer was %v\n", answer)
	}
	
}
