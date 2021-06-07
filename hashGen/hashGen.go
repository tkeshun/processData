package hashGen

import (
	"crypto/sha256"
)

func Generator(readtxt string) []byte {
	hash256 := sha256.Sum256([]byte(readtxt)) 
	return hash256[:]//[32] byte型にならないようにする
}