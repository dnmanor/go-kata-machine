package main

import "math"

func TwoCrystalBalls(storey []bool) int {
	lengthOfBuilding := len(storey)
	jumpLength := int(math.Sqrt(float64(lengthOfBuilding)))
	i := jumpLength

	for i < lengthOfBuilding && !storey[i] {
		i += jumpLength
	}

	i -= jumpLength

	for j := 0; j < jumpLength && i < lengthOfBuilding; j, i = j+1, i+1 {
		if storey[i] {
			return i
		}
	}

	return -1
}
