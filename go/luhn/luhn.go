package luhn

import (
	"fmt"
	"log"
	"unicode"
)

func Valid(id string) bool {
	id, err := checkAndStrip(id)
	if err != nil {
		log.Printf("incorrect number: %s", err)
		return false
	}

	valid := validate(id)
	if !valid {
		return false
	}

	return true
}

func checkAndStrip(id string) (string, error) {
	var resRunes []rune
	for _, r := range []rune(id) {
		if r == ' ' {
			continue
		}
		if unicode.IsDigit(r) {
			resRunes = append(resRunes, r)
			continue
		}
		return string(r), fmt.Errorf("invalid character: %c", r)
	}

	return string(resRunes), nil
}

func validate(id string) bool {
	if len(id) <= 1 {
		return false
	}

	// double every 2nd digit starting from the right, subtract 9 if greater than 9
	idRunes := []rune(id)
	idLen := len(idRunes)
	for i := 0; i < idLen; i++ {
		currentIndex := idLen - i - 1
		if i%2 != 0 {
			d := (int(idRunes[currentIndex]) - '0') * 2
			if d > 9 {
				d -= 9
			}
			idRunes[currentIndex] = []rune(fmt.Sprint(d))[0]
		}
	}

	var sum int
	for _, d := range idRunes {
		sum += (int(d) - '0')
	}

	return sum%10 == 0
}
