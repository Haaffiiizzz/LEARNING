package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to this CLI guessing game. Goodluck!")
	rand.Seed(time.Now().UnixNano())

    randomNumber := rand.Intn(11)

	var guess int
	fmt.Println("\nType in your guess:")

	fmt.Scan(&guess)

	for guess != randomNumber {
		if guess < randomNumber {
			fmt.Println("Guess is lower than the number. Try again!")
			fmt.Scan(&guess)
		} else {
			fmt.Println("Guess is higher than the number. Try again!")
			fmt.Scan(&guess)
		}

	}

	fmt.Println("Congrats, you got it right!")



}
