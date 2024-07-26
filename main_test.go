package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name string
		n    int
		want bool
		msg  string
	}{
		{"prime number", 7, true, "7 is a prime number"},
		{"not prime number", 8, false, "8 is not a prime number because its divisible by 2"},
		{"zero", 0, false, "0 is not prime by definition"},
		{"one", 1, false, "1 is not prime by definition"},
		{"negative number", -1, false, "Negative numbers are not prime by definition"},
	}

	for _, e := range primeTests {
		got, msg := isPrime(e.n)
		// if we expect a prime number but got false
		if e.want && !got {
			t.Errorf("%s failed: got %t, want %t", e.name, got, e.want)
		}

		// if we expect a non-prime number but got true
		if !e.want && got {
			t.Errorf("%s failed: got %t, want %t", e.name, got, e.want)
		}

		if e.msg != msg {
			t.Errorf("%s failed: got %s, want %s", e.name, msg, e.msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	// save the original stdout
	oldOut := os.Stdout

	// create a read and write pipe
	// the pipe will allow us to simulate user input on terminal by writing to the write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close the write pipe
	w.Close()

	// reset os.Stdout to the original stdout
	os.Stdout = oldOut

	// ler o output da nossa func prompt() a partir do read pipe
	// e comparar com o que esperamos
	out, _ := io.ReadAll(r)
	if string(out) != "Enter a number: " {
		t.Errorf("prompt() failed: got %s, want %s", out, "Enter a number: ")
	}
}

func Test_intro(t *testing.T) {
	// save the original stdout
	oldOut := os.Stdout

	// create a read and write pipe
	// the pipe will allow us to simulate user input on terminal by writing to the write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	// close the write pipe
	w.Close()

	// reset os.Stdout to the original stdout
	os.Stdout = oldOut

	// testes
	out, _ := io.ReadAll(r)
	if !strings.Contains(string(out), "Welcome to the prime number checker") {
		t.Errorf("intro() failed: got %s, want %s", out, "Welcome to the prime number checker")
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty", input: "", want: "Please enter a valid number"},
		{name: "zero", input: "0", want: "0 is not prime by definition"},
		{name: "one", input: "1", want: "1 is not prime by definition"},
		{name: "two", input: "2", want: "2 is a prime number"},
		{name: "negative number", input: "-1", want: "Negative numbers are not prime by definition"},
		{name: "typed", input: "three", want: "Please enter a valid number"},
		{name: "decimal", input: "3.14", want: "Please enter a valid number"},
		{name: "quit", input: "q", want: ""},
		{name: "QUIT", input: "Q", want: ""},
		{name: "greek", input: "ΔΟΥΛΕΙΑ", want: "Please enter a valid number"},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		got, _ := checkNumbers(reader)

		// teste
		if !strings.EqualFold(got, e.want) {
			t.Errorf("%s failed: got %s, want %s", e.name, got, e.want)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	// to test this function, we need a channel and an instance of an io.Reader
	doneChan := make(chan bool)

	var stdin bytes.Buffer

	// simulate 1 + enter + q + enter
	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
