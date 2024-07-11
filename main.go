package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var attemps = 0

func main() {
	showMenu()
	randomNumber := generateRandomNumber()
	for {
		guessesNum, err := getNumberFromUser()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		checkGuessedNumber(randomNumber, guessesNum)
	}

}

func generateRandomNumber() int {
	rNum := rand.Intn(100)
	if rNum == 0 {
		return 1
	}
	return rNum
}

func getNumberFromUser() (int, error) {
	setAttemps()
	fmt.Println("\nEnter integer number")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, errors.New("Invalid number, please enter only integer number.")
	}
	return num, nil
}

func checkGuessedNumber(randNum, guessNum int) {
	if randNum == guessNum {
		fmt.Println("\nyou guessed it correct.")
		fmt.Printf("Numbers of attemps took to guess number is %d\n", attemps)
		os.Exit(0)
	}

	if guessNum >= (randNum-3) && guessNum <= (randNum+3) {
		fmt.Println("\nyou are almost there")
	} else if guessNum > randNum {
		fmt.Println("\nguessed number is too high then generate number")
	} else if guessNum < randNum {
		fmt.Println("\nguessed number is too low then generate number")
	}

	fmt.Println("\nLet's guess again")
}

// keeps track of how many attemps the user has made
func setAttemps() {
	attemps += 1
}

func showMenu() {
	fmt.Println("Guess the Number")
	fmt.Println("Guess the number between 1-100")
}
