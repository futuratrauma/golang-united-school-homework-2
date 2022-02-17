package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

type Figure int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesTriangle Figure = 3
	SidesSquare   Figure = 4
	SidesCircle   Figure = 0
)

func CalcSquare(sideLen float64, sidesNum Figure) float64 {
	if sidesNum == SidesTriangle {
		return math.Sqrt(math.Pow(sideLen, 2)-math.Pow(sideLen/2, 2)) * sideLen / 2
	} else if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2)
	} else if sidesNum == SidesCircle {
		return math.Pow(sideLen, 2) * math.Pi
	} else {
		return 0
	}
}
