package utils

func JaroWinkler(s1, s2 string) (similarity float64) {

	// index by code point, not byte
	r1 := []rune(s1)
	r2 := []rune(s2)

	r1Length := len(r1)
	r2Length := len(r2)

	if r1Length == 0 || r2Length == 0 {
		return
	}

	minLength := 0
	if r1Length > r2Length {
		minLength = r1Length
	} else {
		minLength = r2Length
	}

	searchRange := minLength
	searchRange = (searchRange / 2) - 1
	if searchRange < 0 {
		searchRange = 0
	}
	var lowLim, hiLim, transCount, commonChars int
	var i, j, k int

	r1Flag := make([]bool, r1Length+1)
	r2Flag := make([]bool, r2Length+1)

	// find the common chars within the acceptable range
	commonChars = 0
	for i = range r1 {
		if i >= searchRange {
			lowLim = i - searchRange
		} else {
			lowLim = 0
		}

		if (i + searchRange) <= (r2Length - 1) {
			hiLim = i + searchRange
		} else {
			hiLim = r2Length - 1
		}

		for j := lowLim; j <= hiLim; j++ {
			if !r2Flag[j] && r2[j] == r1[i] {
				r2Flag[j] = true
				r1Flag[i] = true
				commonChars++

				break
			}
		}
	}

	// if we have nothing in common at this point, nothing else can be done
	if commonChars == 0 {
		return
	}

	// otherwise we count the transpositions
	k = 0
	transCount = 0
	for i := range r1 {
		if r1Flag[i] {
			for j = k; j < r2Length; j++ {
				if r2Flag[j] {
					k = j + 1
					break
				}
			}
			if r1[i] != r2[j] {
				transCount++
			}
		}
	}
	transCount /= 2

	// adjust for similarities in nonmatched characters
	similarity = float64(commonChars)/float64(r1Length) +
		float64(commonChars)/float64(r2Length) +
		(float64(commonChars-transCount))/float64(commonChars)
	similarity /= 3.0

	return
}
