package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.TrimSpace(input)
	if len(input) <= 0 {
		return "", fmt.Errorf("Empty input: %w", errorEmptyInput)
	}

	elements := []string{}

	elem := ""
	for _, r := range input {
		if string(r) == " " {
			continue
		} else if string(r) == "+" || string(r) == "-" {
			if len(elem) > 0 {
				if (elem != "+") && (elem != "-") {
					elements = append(elements, elem)
				}
				elem = ""
			}
		}
		elem += string(r)
	}

	if len(elem) > 0 {
		if elem != "+" && elem != "-" {
			elements = append(elements, elem)
		}
	}

	if (len(elements) > 2) || (len(elements) <= 1) {
		return "", fmt.Errorf("Wrong amount of operands: %w", errorNotTwoOperands)
	}

	sum := 0
	for _, r := range elements {
		n, err := strconv.Atoi(r)
		if err != nil {
			return "", fmt.Errorf("bad token  %w", err)
		}
		sum += n
	}

	res := strconv.Itoa(sum)
	return res, nil
}
