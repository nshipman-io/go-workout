package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func guessing_game() {
	var userNum int
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(100-0) + 0
	guesses := 3
	for guesses > 0 {
		fmt.Printf("You have %d guesses remaining.\n", guesses)
		fmt.Print("Guess a number between 0 and 100: ")
		_, err := fmt.Scanf("%d", &userNum)

		if err != nil {
			fmt.Println(err)
		}

		if userNum < randomNum {
			fmt.Println("Too low")
			guesses -= 1
		} else if userNum > randomNum {
			fmt.Println("Too high")
			guesses -= 1
		} else {
			fmt.Println("Just right")
			break
		}

	}

	if guesses == 0 {
		fmt.Printf("You ran out of guesses.\n The number was %d. Try again!", randomNum)
		os.Exit(0)
	}
}

func main() {
	guessing_game()
	fmt.Println("Exiting game. Congrats!")
	os.Exit(0)
}
