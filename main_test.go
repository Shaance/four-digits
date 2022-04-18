package main

import (
	"fmt"
	"testing"
)

const lenError = "[Error] Input needs to be 4 characters long"
const notNumericError = "[Error] Input needs to be composed of 4 digits between 0 and 9"
const duplicateError = "[Error] Input can't have duplicates"

func TestValidateInput_TooShort(t *testing.T) {
	testValidateInput("123", lenError, t)
}

func TestValidateInput_TooLong(t *testing.T) {
	testValidateInput("12323", lenError, t)
}

func TestValidateInput_NotNumeric(t *testing.T) {
	testValidateInput("123s", notNumericError, t)
}

func TestValidateInput_Duplicates(t *testing.T) {
	testValidateInput("1232", duplicateError, t)
}

func TestValidateInput_HappyCase(t *testing.T) {
	actual := validateInput("1234")
	if actual != nil {
		t.Errorf("Expected no error, instead got %s", actual)
	}
}

func testValidateInput(input string, errMsg string, t *testing.T) {
	actual := validateInput(input)
	if fmt.Sprint(actual) != errMsg {
		t.Errorf("Expected error to be %s, got %s", errMsg, actual)
	}
}
