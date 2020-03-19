package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
)

var helpstring = `
Call this application on the
file you wish to SHA256 checksum. 

Usage:
    checksum.exe <filename>
`

func main() {
	var fileName string
	if len(os.Args) < 2 {
		//os.Stdout.WriteString(helpstring)
		fileName = waitForUserInput("SHA256\nWrite name of file to be checksum:\n")
	} else {
        if os.Args[1] == "help" || os.Args[1] == "h" || os.Args[1] == "" {
            _, _ = os.Stdout.WriteString(helpstring)
            waitForUserInput("")
            return
        }
		fileName = os.Args[1]
	}

	f, err := os.Open(fileName)
	if err != nil {
		waitForUserInput(err.Error())
		return
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	os.Stdout.WriteString(hex.EncodeToString(h.Sum(nil)))
	waitForUserInput("")
}

func waitForUserInput(message string) string {
    if message == "" {
        message = "\nPress a [Enter] to continue...\n"
    }
	input := bufio.NewScanner(os.Stdin)
	os.Stdout.WriteString(message)
	input.Scan()
	return input.Text()
}
