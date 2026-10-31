package main

import (
	"io"
	"os"
	"testing"
)

func TestIsPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"zero", 0, false, "0 is not prime, by definition"},
		{"one", 1, false, "1 is not prime, by definition"},
		{"prime", 2, true, "2 is a prime number"},
		{"composite", 4, false, "4 is not a prime because it is divisible by 2"},
		{"prime", 7, true, "7 is a prime number"},
		{"composite", 9, false, "9 is not a prime because it is divisible by 3"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s : expected true but got false", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s : expected false but got true", e.name)
		}
		if e.msg != msg {
			t.Errorf("%s : expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	//save a copy  of os.stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	//set os.stdout to our write pipe
	os.Stdout = w

	prompt()
	//close our writer
	_ = w.Close()

	// reset os.stdout to waht it was before
	os.Stdout = oldOut

	// read the out for our prompt() func for our read pipe
	out, _ := io.ReadAll(r)

	//perform our test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt : expected -> but  go %s", string(out))
	}

}
