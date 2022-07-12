package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		// read keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// empty input is skipped as in other shells
		if input == "" {
			continue
		}

		// handle result of executing the input
		if err = execute(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execute(input string) error {
	// remove newline
	input = strings.TrimSuffix(input, "\n")

	// separate commands and args on spaces
	args := strings.Split(input, " ")

	// check for built-in commands that need special handling
	switch args[0] {
	case "cd $arg":
		// cd to home dir with empty path not supported
		if len(args) < 1 {
			return errors.New("A specific path is required")
		}
		// cd into specified path held in args[1] and return the error
		return os.Chdir(args[1])

	case "echo":
		// the inputs that follow args[0] are echoed back
		var str, temp string
		for i := 1; i < len(os.Args); i++ {
			str += temp + os.Args[i]
			temp = " "
		}
		fmt.Println(str)

	case "exit":
		// does what is says
		os.Exit(0)
	}

	// prepare whichever command that will be executed
	cmd := exec.Command(args[0], args[1:]...)

	// set output
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// execute command and return the error
	return cmd.Run()
}
