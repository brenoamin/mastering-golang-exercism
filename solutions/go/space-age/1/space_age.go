package spaceage

import "math"

type Planet string


func Age(seconds float64, planet Planet) float64 {
	base := float64(31557600)
    period := planetsOrbitalPeriodInEarthDays(planet)

    if period < 0 {
        return period
    }
    
	conversion := (seconds / base) / period

	return math.Round(conversion*100) /100
}

func planetsOrbitalPeriodInEarthDays(planet Planet) float64 {
	periods := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

   value, ok := periods[planet]
	if !ok {
		return -1
	}
	return value

}
