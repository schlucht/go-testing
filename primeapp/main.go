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
	// scrivi un messaggio di benvenuto
	intro()

	// creare un canale per indicare quando l'utente vuole uscire
	// Erstellen Sie einen Kanal, um anzuzeigen, wann der Benutzer aufhören möchte
	doneChan := make(chan bool)

	// start a goroutine to read user input and run program
	go readUserInput(os.Stdin, doneChan)
	// block until the doneChan gets a value
	<-doneChan
	// close the channel
	close(doneChan)
	// say goodby
	fmt.Println("Tschüss")

}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumber(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumber(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	numToChekc, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "please enter a whole number!", false
	}

	_, msg := isPrime(numToChekc)

	return msg, false

}

func intro() {
	s := "Willkommen zum Primzahlen berechner"
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("\t%s\n", s)
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Gib eine ganze Zahl ein und wir sagen ob eseine Primzahl ist. Mit q beenden")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func isPrime(n int) (bool, string) {
	// 0 e 1 sono per definizione non primi
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d non è numero primo!", n)
	}

	if n < 0 {
		return false, "I numeri negativi non sono numeri primi"
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d non è numero primo perchè divisibile per %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d è numero primo!", n)
}
