package robinkarp

import (
	"testing"
)

func TestMatch(t *testing.T) {
	rk := NewRobinKarp("needle")
	
	text := "There is a needle in the haystack"
	textBytes := []byte(text)

	num := 0
	for _, c := range textBytes {
		if gotit := rk.SearchNextChar(c); gotit == true {
			num++
		}
	}

	if !(num == 1) {
		t.Error("Could not get needle in the haystack when there was a needle present")
	}

	if (num > 1) {
		t.Error("Found more than one needle when there was only one.")
	}
}

func TestNoMatch(t *testing.T) {
	rk := NewRobinKarp("needle")
	
	text := "There is not n eedle in the haystack"
	textBytes := []byte(text)

	num := 0
	for _, c := range textBytes {
		if gotit := rk.SearchNextChar(c); gotit == true {
			num++
		}
	}

	if !(num == 0) {
		t.Error("Got a needle in the haystack when there was NO needle present")
	}

}

func TestMultMatch(t *testing.T) {

	rk := NewRobinKarp("needle")
	
	text := "There are two needles in the needle haystack"
	textBytes := []byte(text)

	num := 0
	for _, c := range textBytes {
		if gotit := rk.SearchNextChar(c); gotit == true {
			num++
		}
	}

	if !(num == 2) {
		t.Error("Got a different number than 2 when there are two needles in the haystack.")
	}

	for _, c := range textBytes {
		if gotit := rk.SearchNextChar(c); gotit == true {
			num++
		}
	}

	if !(num == 4) {
		t.Error("Expecting 4 needles when I search twice in the haystack.")
	}

}