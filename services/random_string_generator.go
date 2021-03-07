package services

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomStringGenerator(n int) string {
	b := make([]rune, n)
	lettersLength := int64(len(letters))

	for i := range b {
		b[i] = letters[rand.Int63()%lettersLength]
	}

	return string(b)
}
