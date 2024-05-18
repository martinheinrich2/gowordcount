package main

import (
	"flag"
	"fmt"
	"github.com/martinheinrich2/gowordcount/cmd"
)

func main() {
	// Define the command line options
	var arg string
	var filename string
	var help bool

	// Set the reference to the variable, name of the argument/flag, default value and help text
	flag.StringVar(&arg, "flag", "w", "Flags: b, l, w, c")
	flag.StringVar(&filename, "file", "test.txt", "Filename")
	flag.BoolVar(&help, "help", false, "Show help")
	flag.Parse()

	if help {
		flag.Usage()
		flag.PrintDefaults()
	}

	switch arg {
	case "b":
		fmt.Println("Number of bytes:", cmd.CountBytes(filename))
	case "l":
		fmt.Println("Number of lines:", cmd.CountLines(filename))
	case "w":
		fmt.Println("Number of words:", cmd.CountWords(filename))
	case "c":
		fmt.Println("Number of Characters:", cmd.CountCharacters(filename))
	}
}
