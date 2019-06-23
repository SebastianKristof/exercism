/*
Package space is used for calculating planet-years and other useful information
about the solar system and its inhabitants.
*/
package space

import "log"

// Planet is the type that holds the planet name as a string value
type Planet string

// Age calculates the age in planet-years when given
// the age in seconds and the planet
func Age(seconds float64, planet Planet) float64 {

	if seconds < 0 {
		log.Fatalf("negative age doesn't make sense")
	}

	earthOrbitalPeriodSeconds := 31557600.00
	ageInEarthYears := seconds / earthOrbitalPeriodSeconds

	yearLengthRatios := map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	yearLengthRatio, ok := yearLengthRatios[planet]
	if !ok {
		log.Fatalf("wrong planet name")
	}

	return ageInEarthYears / yearLengthRatio
}
