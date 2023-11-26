package util

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	shortCodeLength   = 8
	lettersAndNumbers = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generateRandomCode() string {
	result := make([]byte, shortCodeLength)
	for i := range result {
		result[i] = lettersAndNumbers[rand.Intn(len(lettersAndNumbers))]
	}
	return string(result)
}

func GenerateUniqueCode(checkUnique func(string) bool) (string, error) {
	rand.Seed(time.Now().UnixNano())

	for attempt := 0; attempt < 3; attempt++ {
		shortCode := generateRandomCode()
		if !checkUnique(shortCode) {
			continue
		}

		return shortCode, nil
	}
	return "", fmt.Errorf("failed to generate a unique short code after 3 attempts")
}
