package patternmatcher

// PatternsAreEqual compares two patterns to determine if they are equal.
// It checks if the patterns have the same number of coordinates and if each
// corresponding coordinate in both patterns is equal. If the patterns are
// equal, it returns true; otherwise, it returns false.
func PatternsAreEqual(p1, p2 Pattern) bool {
	if len(p1) != len(p2) {
		return false
	}
	for i := range p1 {
		if p1[i] != p2[i] {
			return false
		}
	}
	return true
}
