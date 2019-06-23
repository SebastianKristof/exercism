/*
Package raindrops is useful for understanding the rain.
*/
package raindrops

import "fmt"

type raindrop struct {
	factor int
	sound  string
}

// Convert translates the number of drops into rainspeak.
func Convert(num int) string {

	raindrops := []raindrop{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}

	rainSays := ""

	// if factors are found, grow the string
	for _, drop := range raindrops {
		if num%drop.factor == 0 {
			rainSays += drop.sound
		}
	}

	// if rain has said nothing so far, just speak the number
	if rainSays == "" {
		// could also use strconv.Itoa() here
		rainSays = fmt.Sprint(num)
	}

	return rainSays
}
