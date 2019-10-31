package space

// Planet is a placeholder for which planet the user inputs
type Planet string

// EarthYear is how many seconds it takes the earth to revolve around the sun once
const EarthYear = 31557600

// Age sets up the map/dictionary for the planets. It assigns each planet to a decimal of how many seconds its own revolution takes
func Age(days float64, planet Planet) float64 {
	PlanetYrs := map[Planet]float64{
		"Mercury": EarthYear * 0.2408467,
		"Venus":   EarthYear * 0.61519726,
		"Earth":   EarthYear,
		"Mars":    EarthYear * 1.8808158,
		"Jupiter": EarthYear * 11.862615,
		"Saturn":  EarthYear * 29.447498,
		"Uranus":  EarthYear * 84.016846,
		"Neptune": EarthYear * 164.79132,
	}
	return PlanetAgeCalc(PlanetYrs[planet], days)
}

// PlanetAgeCalc is used to calculate the years on the given planet based on the value passed by the dictionary PlanetYrs
func PlanetAgeCalc(planetYr float64, seconds float64) float64 {
	return seconds / planetYr
}
