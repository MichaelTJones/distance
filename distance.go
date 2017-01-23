package distance

// Jaro and Jaro-Winkler string distance (similarity) metrics
// good background reading
//   http://www.cs.cmu.edu/~wcohen/postscript/ijcai-ws-2003.pdf

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Compute Jaro similarity between two strings
func Jaro(a, b string) float64 {
	// handle easy cases
	switch {
	case a == b:
		return 1.0
	case a == "" || b == "":
		return 0.0
	}

	// ensure a is not longer than b
	// if len(a) > len(b) {
	// 	a, b = b, a
	// }

	lenA := len(a)
	lenB := len(b)
	window := max(lenA, lenB)/2 - 1

	// find and record matches
	match := 0
	am := make([]bool, lenA)
	bm := make([]bool, lenB)
	for ia := range a {
		left := max(0, ia-window)
		right := min(ia+window+1, lenB)
		for ib := left; ib < right; ib++ {
			if a[ia] == b[ib] && !bm[ib] {
				am[ia] = true
				bm[ib] = true
				match++
				break
			}
		}
	}

	score := 0.0
	transpose := 0
	if match != 0 {
		// count transpositions
		left := 0
		for ia := range a {
			if am[ia] {
				ib := left
				for ; ib < lenB; ib++ {
					if bm[ib] {
						left = ib + 1
						break
					}
				}
				if a[ia] != b[ib] {
					transpose++
				}
			}
		}
		transpose >>= 1

		// compute score by collecting terms to minimize rounding error
		score = float64(match*match*(lenA+lenB)+lenA*lenB*(match-transpose)) / float64(3*lenA*lenB*match)
	}

	return score
}

// Compute Jaro–Winkler similarity between two strings
func JaroWinkler(a, b string) float64 {
	limit := min(4, min(len(a), len(b)))
	prefix := 0
	for i := 0; i < limit && a[i] == b[i]; i++ {
		prefix++
	}

	jaro := Jaro(a, b)
	weight := 0.1
	return jaro + weight*float64(prefix)*(1.0-jaro)
}
