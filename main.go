package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a welcome message
	intro()

	// create a channel
	doneChan := make(chan bool)

	// start a goroutine to read user input and run program
	go readUserInput(os.Stdin, doneChan)

	// block until the program is done ( the channel gets a value )
	<-doneChan

	// close the channel
	close(doneChan)

	// say goodbye
	fmt.Println("Goodbye!")
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime numbers
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime by definition", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime by definition"
	}

	// we can use the mod operator repeatedly to see if we have a prime number or not
	// we keep until n/2 because a number is not divisible by any number greater than half of it
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// if no remainder not a prime number
			return false, fmt.Sprintf("%d is not a prime number because its divisible by %d", n, i)
		}
	}

	// mostly true
	return true, fmt.Sprintf("%d is a prime number", n)
}

func intro() {
	fmt.Println("Welcome to the prime number checker")
	fmt.Println("-----------------------------------")
	fmt.Println("Enter a number to check if it is a prime number")
	fmt.Println("Enter 'q' to quit the program")
	prompt()
}

func prompt() {
	fmt.Print("Enter a number: ")
}

func readUserInput(in io.Reader, doneCh chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)
		if done {
			doneCh <- true
			return
		}

		// print the result
		fmt.Println(res)
		// prompt the user again
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()

	// check to see if the user wants to quit
	// the Equalfold method is case insensitive, it will return true if the strings are equal
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convert what the user entered to an integer
	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a valid number", false
	}

	_, msg := isPrime(numToCheck)
	return msg, false
}
