package services

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomStringGenerator(n int) string {
	b := make([]rune, n)
	lettersLength := len(letters)

	for i := range b {
		b[i] = letters[rand.Intn(lettersLength)]
	}

	return string(b)
}
