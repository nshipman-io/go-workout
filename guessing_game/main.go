package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func guessing_game() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)
	randomNum := rand.Intn(100-0) + 0
	guesses := 3
	for true {
		if guesses == 0 {
			fmt.Printf("You ran out of guesses.\n The number was %d. Try again!", randomNum)
			os.Exit(0)
		}

		fmt.Printf("You have %d guesses remaining.\n", guesses)
		fmt.Print("Guess a number between 0 and 100: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("An error has occured while reading input.")
		} else {
			input = strings.TrimSuffix(input, "\n")
			userNum, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input. Please enter an integer value.")
			} else {
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
		}
	}

}

func main() {
	guessing_game()
	fmt.Println("Exiting game. Congrats!")
	os.Exit(0)
}
