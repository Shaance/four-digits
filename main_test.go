package main

import (
	"testing"
	"fmt"
)

const len_error = "[Error] Input needs to be 4 characters long"
const not_numeric_error = "[Error] Input needs to be composed of 4 digits between 0 and 9" 
const duplicate_error = "[Error] Input can't have duplicates"

func TestValidateInput_TooShort(t *testing.T) {
	testValidateInput("123", len_error, t)
}

func TestValidateInput_TooLong(t *testing.T) {
	testValidateInput("12323", len_error, t)
}

func TestValidateInput_NotNumeric(t *testing.T) {
	testValidateInput("123s", not_numeric_error, t)
}

func TestValidateInput_Duplicates(t *testing.T) {
	testValidateInput("1232", duplicate_error, t)
}

func TestValidateInput_HappyCase(t *testing.T) {
	actual := validate_input("1234")
	if actual != nil {
		t.Errorf("Expected no error, instead got %s", actual)
	}
}

func testValidateInput(input string, err_msg string, t *testing.T) {
	actual := validate_input(input)
	if fmt.Sprint(actual) != err_msg {
		t.Errorf("Expected error to be %s, got %s", err_msg, actual)
	}
}