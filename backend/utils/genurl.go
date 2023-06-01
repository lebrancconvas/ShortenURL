package utils

import (
	"math/rand"
	"time"
)

func GenerateShortURL() string {
	const alphabetLowercase = "abcdefghijklmnopqrstuvwxyz"
	const alphabetUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const number = "0123456789"
	const charSet = alphabetLowercase + alphabetUppercase + number

	rand.Seed(time.Now().UnixNano())

	var shortURL string
	for i := 0; i < 6; i++ {
		shortURL += string(charSet[rand.Intn(len(charSet))])
	}

	return shortURL
}
