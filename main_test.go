package main

import (
	"fmt"
	"testing"
)

const lenError = "[Error] Input needs to be 4 characters long"
const notNumericError = "[Error] Input needs to be composed of 4 digits between 0 and 9"
const duplicateError = "[Error] Input can't have duplicates"

func TestValidateInput_TooShort(t *testing.T) {
	testValidate(createUserInputWithStringInput("123"), lenError, t)
}

func TestValidateInput_TooLong(t *testing.T) {
	testValidate(createUserInputWithStringInput("12345"), lenError, t)
}

func TestValidateInput_NotNumeric(t *testing.T) {
	testValidate(createUserInputWithStringInput("123s"), notNumericError, t)
}

func TestValidateInput_Duplicates(t *testing.T) {
	testValidate(createUserInputWithStringInput("1232"), duplicateError, t)
}

func TestValidateInput_HappyCase(t *testing.T) {
	userInput := createUserInputWithStringInput("1234")
	userInput.validate()
	if userInput.err != nil {
		t.Errorf("Expected no error, instead got %s", userInput.err.Error())	
	}
}

func createUserInputWithStringInput(stringInput string) UserInput {
	var userInput UserInput
	userInput.originalStringInput = stringInput
	return userInput
}
 
func testValidate(input UserInput, expectedErrMsg string, t *testing.T) {
	input.validate()
	actualErr := input.err.Error()
	if fmt.Sprint(actualErr) != expectedErrMsg {
		t.Errorf("Expected error to be %s, got %s", expectedErrMsg, actualErr)
	}
}
