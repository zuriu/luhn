// Copyright 2013 DigiExam
// Licensed under MIT

// Package to generate and validate luhn numbers
// http://en.wikipedia.org/wiki/Luhn_algorithm

package luhn

import (
	"fmt"
	"strconv"
)

// Calculates the luhn on the number string
func Luhn(numbers string) (int, error) {
	if _, err := strconv.ParseInt(numbers, 10, 64); err != nil {
		return 0, err
	}

	checksum 	:= 0
	result		:= 0

	// Calculate the checksum, from right to left.
	for i := len(numbers) - 1; i >= 0; i-- {
		result = int(numbers[i] - '0')
		isOdd := (len(numbers) - i + 1) % 2 == 0

		if isOdd {
			result *= 2

			if 10 <= result {
				result -= 9
			}
		}

		checksum += result			
	}

	checksum = (checksum % 10)

	if checksum == 0 {
		return 0, nil
	} else {
		return 10 - checksum, nil
	}
}

// Returns the string with the luhn number appended
func Append(numbers string) (string, error) {
	if l, err := Luhn(numbers); err == nil {
		return fmt.Sprintf("%v%v", numbers, l), nil
	} else {
		return "", err
	}
}

// Validates that the string got a valid luhn number
// If one would like to optmize the luhn validation you could change the algorithm to: http://orb-of-knowledge.blogspot.se/2009/08/extremely-fast-luhn-function-for-c.html
func IsValid(numbers string) (bool, error) {
	numberPart 	:= numbers[0:len(numbers)-1]
	luhnPart 	:= numbers[len(numbers)-1:]

	l, err := Luhn(numberPart); 
	if err == nil {
		if lp, err := strconv.Atoi(luhnPart); err == nil {
			isValid := l == lp
			return isValid, nil
		} else {
			return false, err
		}
	} else {
		return false, err		
	}
}