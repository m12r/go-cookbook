package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	// Get file descriptor of stdin.
	fd := int(os.Stdin.Fd())

	// Check if stdin is connected to a terminal.
	// Stdin is not a terminal, if data is piped to this program.
	// Try: `echo "Hello" | password_from_stdin`.
	// You will notice, that the program does not get past this code block.
	if !term.IsTerminal(fd) {
		fmt.Fprintln(os.Stderr, "error: stdin not a terminal")
		os.Exit(1)
	}

	// Safe the state of the terminal and reset it if the program is finished.
	// If you don't do this, you will set the terminal, so it won't echo stdin anymore.
	safeState, err := term.GetState(fd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: cannot safe stdin state: %v", err)
		os.Exit(1)
	}
	defer term.Restore(fd, safeState)

	// Show prompt and read the password from stdin.
	fmt.Print("Enter password: ")
	password, err := term.ReadPassword(fd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: cannot read password from stdin: %v", err)
		os.Exit(1)
	}

	// Print the read password to stdout.
	fmt.Printf("Your password is: %q\n", string(password))
}
