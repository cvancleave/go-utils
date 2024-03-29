package utils

func LevenshteinDistance(s1, s2 string) (distance int) {

	r1 := []rune(s1)
	r2 := []rune(s2)

	r1Length := len(r1)
	r2Length := len(r2)

	if r1Length == 0 || r2Length == 0 {
		return
	}

	vec1 := make([]int, len(s2)+1)
	vec2 := make([]int, len(s2)+1)

	// initializing vec1
	for i := 0; i < len(vec1); i++ {
		vec1[i] = i
	}

	// initializing the matrix
	for i := 0; i < len(r1); i++ {
		vec2[0] = i + 1

		for j := 0; j < len(r2); j++ {
			cost := 1
			if r1[i] == r2[j] {
				cost = 0
			}
			min := minimum(vec2[j]+1,
				vec1[j+1]+1,
				vec1[j]+cost)
			vec2[j+1] = min
		}

		for j := 0; j < len(vec1); j++ {
			vec1[j] = vec2[j]
		}
	}

	return vec2[len(r2)]
}

// Helper function for Levenshtein distance
func minimum(value0 int, values ...int) int {
	min := value0
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}
