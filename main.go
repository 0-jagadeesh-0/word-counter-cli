package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	COUNT      string = "count"
	BYTES      string = "bytes"
	WORDS      string = "words"
	LINES      string = "lines"
	CHARACTERS string = "chars"
	EXIT       string = "exit"
	HELP       string = "help"
)

func ToolInfo() {
	fmt.Println(" ------------------------------------------------------------------------------------------------------------------ ")
	fmt.Println("| COMMAND                     DESCRIPTION                                                                          |")
	fmt.Println(" ------------------------------------------------------------------------------------------------------------------ ")
	fmt.Println("| count bytes 'file.txt'      Counts total bytes in the file.(Replace 'file.txt' with path of the text file.)      |")
	fmt.Println("| count words 'file.txt'      Counts total words in the file.(Replace 'file.txt' with path of the text file.)      |")
	fmt.Println("| count lines 'file.txt'      Counts total lines in the file.(Replace 'file.txt' with path of the text file.)      |")
	fmt.Println("| count chars 'file.txt'      Counts total characters in the file.(Replace 'file.txt' with path of the text file.) |")
	fmt.Println("| exit                        To exit                                                                              |")
	fmt.Println(" ------------------------------------------------------------------------------------------------------------------ ")
}

func CountBytes(data string) int {
	return len(data)
}

func CountCharacters(data string) int {
	return utf8.RuneCountInString(data)
}

func CountLines(data string) int {
	lines := strings.Split(data, "\n")
	lineCount := len((lines))
	if lineCount > 0 && lines[lineCount-1] == "" {
		return lineCount - 1
	}
	return lineCount
}

func CountWords(data string) int {
	words := strings.Fields(data)
	return len(words)
}

func InvalidCommandError() {
	fmt.Println("Command not recognized. Please enter 'help' to view available commands.")
}

func ValidateCommand(commandArgs []string) bool {
	if len(commandArgs) == 0 {
		return false
	} else if len(commandArgs) > 3 {
		InvalidCommandError()
		return false
	} else if firstArg := commandArgs[0]; len(commandArgs) == 1 && firstArg != EXIT && firstArg != HELP {
		InvalidCommandError()
		return false
	} else if len(commandArgs) == 2 {
		InvalidCommandError()
		return false
	} else if len(commandArgs) == 3 {
		switch commandArgs[1] {
		case WORDS, LINES, CHARACTERS, BYTES:
			return true
		default:
			InvalidCommandError()
			return false
		}
	} else {
		return true
	}
}

func Execute(commandArgs []string) {
	if commandArgs[0] == EXIT {
		fmt.Println("See you again...")
		os.Exit(0)
	}
	if commandArgs[0] == HELP {
		ToolInfo()
		return
	}
	if !strings.HasSuffix(commandArgs[2], ".txt") {
		fmt.Println("Invalid file type. Accepts only .txt files.")
		return
	}
	data, err := os.ReadFile(commandArgs[2])
	if err != nil {
		fmt.Println("Error while opening the file: ", err)
		return
	}
	switch commandArgs[1] {
	case BYTES:
		fmt.Printf("The file has %v byte(s).\n", CountBytes(string(data)))
	case LINES:
		fmt.Printf("The file has %v line(s).\n", CountLines(string(data)))
	case WORDS:
		fmt.Printf("The file has %v word(s).\n", CountWords(string(data)))
	case CHARACTERS:
		fmt.Printf("The file has %v character(s).\n", CountCharacters(string(data)))
	default:
		InvalidCommandError()
	}
}

func EnterCommand() string {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for {
		if scanner.Scan() {
			input = scanner.Text()
			return input
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			break
		}
	}
	return input
}

func main() {
	if len(os.Args) > 1 {
		fmt.Println("No extra arguments are accepted. Please run 'word-counter-cli' without any arguments for interactive mode.")
		os.Exit(1)
	}
	fmt.Println("\033[1mWelcome to the word counter CLI tool...\033[0m")
	fmt.Println("For any help. Please enter 'help' in the terminal.")
	for {
		fmt.Print("> ")
		command := EnterCommand()
		commandArgs := strings.Split(command, " ")
		isValidCommand := ValidateCommand(commandArgs)
		if isValidCommand {
			Execute(commandArgs)
		}
	}
}
