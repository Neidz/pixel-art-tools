package patternmatcher

import "testing"

func TestPatternsAreEqual(t *testing.T) {
	pattern1 := Pattern{{0, 0}, {1, 1}, {2, 2}}
	pattern2 := Pattern{{0, 0}, {1, 1}, {2, 2}}
	if !PatternsAreEqual(pattern1, pattern2) {
		t.Errorf("Expected patterns to be equal, but they are not.")
	}

	pattern3 := Pattern{{0, 0}, {1, 1}}
	if PatternsAreEqual(pattern1, pattern3) {
		t.Errorf("Expected patterns to be unequal due to different lengths, but they are considered equal.")
	}

	pattern4 := Pattern{{0, 0}, {1, 1}, {2, 3}}
	if PatternsAreEqual(pattern1, pattern4) {
		t.Errorf("Expected patterns to be unequal due to different coordinates, but they are considered equal.")
	}

	emptyPattern1 := Pattern{}
	emptyPattern2 := Pattern{}
	if !PatternsAreEqual(emptyPattern1, emptyPattern2) {
		t.Errorf("Expected empty patterns to be equal, but they are not.")
	}

	if PatternsAreEqual(emptyPattern1, pattern1) {
		t.Errorf("Expected empty and non-empty patterns to be unequal, but they are considered equal.")
	}
}
