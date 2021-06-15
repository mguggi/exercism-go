package space

type Planet string

// Age calculates how old someone would be on the given planet
func Age(s float64, p Planet) float64 {
	var op float64
	switch p {
	case "Earth":
		op = 1.0
	case "Mercury":
		op = 0.2408467
	case "Venus":
		op = 0.61519726
	case "Mars":
		op = 1.8808158
	case "Jupiter":
		op = 11.862615
	case "Saturn":
		op = 29.447498
	case "Uranus":
		op = 84.016846
	case "Neptune":
		op = 164.79132
	}
	return s / (31557600 * op)
}
