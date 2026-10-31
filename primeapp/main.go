package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a welcome message
	intro()

	// create a channel to indicate when the user wants to quit
	doneChan := make(chan bool)

	go readUserInput(doneChan)

	// block until the doneChan gets a value
	<-doneChan

	fmt.Println("Goodbye")
}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func intro() {
	fmt.Println("Is it Prime")
	fmt.Println("________________")
	fmt.Println("Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	if !scanner.Scan() {
		return "", true
	}

	text := strings.TrimSpace(scanner.Text())

	// check to see if the user wants to quit
	if strings.EqualFold(text, "q") {
		return "", true
	}

	// try to convert what the user typed into an int
	numCheck, err := strconv.Atoi(text)
	if err != nil {
		return "Please enter a whole number", false
	}

	// call isPrime and return the result message
	_, msg := isPrime(numCheck)
	return msg, false
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition", n)
	}

	// Negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition"
	}

	// Use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// Not a prime number
			return false, fmt.Sprintf("%d is not a prime because it is divisible by %d", n, i)
		}
	}

	// If the loop completes without finding divisors, it is prime
	return true, fmt.Sprintf("%d is a prime number", n)
}
