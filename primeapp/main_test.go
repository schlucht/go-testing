package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	// result, msg := isPrime(0)
	// if result {
	// 	t.Errorf("with %d as test parameter, got true, but expected false", 0)
	// }
	// if msg != "0 non è numero primo!" {
	// 	t.Error("wrong message returned:", msg)
	// }
	primeTests := []struct {
		name    string
		testNum int
		exeptec bool
		msg     string
	}{
		{"prime", 7, true, "7 è numero primo!"},
		{"zero", 0, false, "0 non è numero primo!"},
		{"prime", 8, false, "8 non è numero primo perchè divisibile per 2"},
		{"one", 1, false, "1 non è numero primo!"},
		{"negatic number", -1, false, "I numeri negativi non sono numeri primi"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.exeptec && !result {
			t.Errorf("%s: excepted true but got false", e.name)
		}
		if !e.exeptec && result {
			t.Errorf("%s: excepted false but got true", e.name)
		}
		if e.msg != msg {
			t.Errorf("%s: expected %s but go %s", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {

	// speichern einer Kopie von os.Stdout
	oldOut := os.Stdout

	// eine Pipe zum schreiben und lesen erstellen
	r, w, _ := os.Pipe()

	// in den os.Stdout schreiben
	os.Stdout = w

	prompt()

	// writer wieder schliessen
	_ = w.Close()

	// reset was vorher im os.Stdout war
	os.Stdout = oldOut

	// liest den autput von prompt() aus
	out, _ := io.ReadAll(r)

	// vergleicht den output vom Prompt
	if string(out) != "-> " {
		t.Errorf("incorect Prompt: exepted -> but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {

	// speichern einer Kopie von os.Stdout
	oldOut := os.Stdout

	// eine Pipe zum schreiben und lesen erstellen
	r, w, _ := os.Pipe()

	// in den os.Stdout schreiben
	os.Stdout = w
	intro()

	// writer wieder schliessen
	_ = w.Close()

	// reset was vorher im os.Stdout war
	os.Stdout = oldOut

	// liest den autput von prompt() aus
	out, _ := io.ReadAll(r)

	// vergleicht den output vom Prompt
	if !strings.Contains(string(out), "Gib eine ganze Zahl ein") {
		t.Errorf("input Text not correct got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "quit", input: "q", expected: ""},
		{name: "zero", input: "0", expected: ""},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumber(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: exepted is incorrect", e.name)
		}
	}
}
