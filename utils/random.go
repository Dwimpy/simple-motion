package utils

import "math/rand"

var alphabet = "abcdefghijklmnoprstuvwxyz"
var emails = []string{"@gmail.com", "@yahoo.com", "@www.com", "@live.com", "@micro.com"}

func RandomString(length int) string {

	var bits []rune
	alphabet_len := len(alphabet)
	for i := 0; i < length; i++ {
		index := rand.Intn(alphabet_len)
		bits = append(bits, rune(alphabet[index]))
	}
	return string(bits)
}

func RandomEmail() string {
	return RandomString(8) + emails[rand.Intn(5)]
}
