package robinkarp

import (
//	"fmt"
)

type RobinKarp struct {
	patternHash		int64
	patternLength	int
	alphaSize		int
	longPrime		int64
	pattern         []byte
	gotStream		[]byte
	idx				int
	curHash			int64
	rm				int64
}

func hash(barr []byte, alphaSize int, longPrime int64) int64 {
	var h int64
	h = 0
	for _, val := range barr {
		h = (int64(alphaSize) * h + int64(val)) % longPrime
	}
	return h
}

// Return a new object that can be used to search for the pattern given in "pattern"
func NewRobinKarp(pattern string) *RobinKarp {
	patternB := []byte(pattern)
	var longPrime int64
	longPrime = 71993
	alphaSize := 256
	patternLength := len(pattern)

	patternHash := hash(patternB, alphaSize, longPrime)
	
	gotStream := make([]byte, patternLength)
	idx := 0
	
	var rm int64
	rm = 1
	for i := 1; i <= (patternLength - 1); i++ {
		rm = (int64(alphaSize) * rm) % longPrime
	}
	 
	return &RobinKarp{patternHash: patternHash, 
		patternLength: patternLength, 
		alphaSize: alphaSize, 
		longPrime: longPrime, 
		pattern: patternB,
		gotStream: gotStream,
		rm: rm,
		idx: idx}
}

func (rk *RobinKarp) GetPatHash() int64 {
	return rk.patternHash
}

func (rk *RobinKarp) GetCurHash() int64 {
	return rk.curHash
}

// Add the next byte to the search pattern and return true if the pattern 
// matches
func (rk *RobinKarp) SearchNextChar(c byte) bool {

	if rk.idx < rk.patternLength - 1 {
		rk.gotStream[rk.idx] = c
	} else if rk.idx == rk.patternLength - 1 {
		rk.gotStream[rk.idx] = c
		rk.curHash = hash(rk.gotStream, rk.alphaSize, rk.longPrime)
	} else {
		substituteIdx := (rk.idx ) % rk.patternLength
		substituteByte := int64(rk.gotStream[substituteIdx])
		rk.curHash = ((rk.curHash + substituteByte*(rk.longPrime - rk.rm))*int64(rk.alphaSize) + int64(c))%rk.longPrime
		rk.gotStream[substituteIdx] = c
	}

	rk.idx ++
	
	return rk.curHash == rk.patternHash
}