package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
)

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder
	var nextCharSet string = lowerCharSet

	if minSpecialChar > 0 {
		nextCharSet += specialCharSet
	}
	if minNum > 0 {
		nextCharSet += numberSet
	}
	if minUpperCase > 0 {
		nextCharSet += upperCharSet
	}

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase

	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(nextCharSet))
		password.WriteString(string(nextCharSet[random]))
	}

	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}

func main() {
	minSpecialChar := flag.Int("s", 1, "Minimum amount of special characters")
	minNum := flag.Int("n", 1, "Minimum numbers")
	minUpperCase := flag.Int("u", 1, "Minimum uppercased characters")
	passwordLength := flag.Int("l", 24, "Password lenght")
	seed := flag.Int("seed", 1337, "Additional custom seed")

	flag.Parse()
	rand.Seed(time.Now().Unix() + int64(*seed))

	password := generatePassword(*passwordLength, *minSpecialChar, *minNum, *minUpperCase)
	fmt.Println(password)
}
