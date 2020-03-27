package util

import "math/rand"

var alphaNumRunes = []rune("ABCDEFGHIJKLmNOPQRSTUVwXYZ1234578 ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes))]
	}
	return string(b)
}