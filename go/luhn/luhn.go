package luhn

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

func Valid(cleanID string) bool {
	cleanID, err := checkAndStrip(cleanID)
	if err != nil {
		log.Printf("incorrect number: %s", err)
		return false
	}

	valid := validate(cleanID)
	if !valid {
		return false
	}

	return true
}

func checkAndStrip(id string) (string, error) {
	var runesID []rune
	for _, r := range id {
		if r == ' ' {
			continue
		}
		if unicode.IsDigit(r) {
			runesID = append(runesID, r)
			continue
		}
		return "", fmt.Errorf("invalid character: %c", r)
	}

	return string(runesID), nil
}

func validate(cleanID string) bool {
	if len(cleanID) <= 1 {
		return false
	}

	// convert string to slice of integer digits, ignore error because the string is already checked
	numbers, _ := strToIntSlice(cleanID)

	// multiply according to Luhn algorithm
	multiplyForLuhn(numbers)

	var sum int
	for _, d := range numbers {
		sum += d
	}

	return sum%10 == 0
}

func strToIntSlice(id string) ([]int, error) {
	var numbers []int
	for _, char := range id {
		n, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, n)
	}
	return numbers, nil
}

func multiplyForLuhn(numbers []int) {
	// double every 2nd digit starting from the right, subtract 9 if greater than 9
	for i := 0; i < len(numbers); i++ {
		currentIndex := len(numbers) - 1 - i // going from last to first
		if i%2 != 0 {
			d := numbers[currentIndex] * 2
			if d > 9 {
				d -= 9
			}
			numbers[currentIndex] = d
		}
	}
}
