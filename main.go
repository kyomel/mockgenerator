package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var help bool
	var inputPath, outputPath string

	flag.BoolVar(&help, "h", false, "Show help")
	flag.BoolVar(&help, "help", false, "Show help")

	flag.StringVar(&inputPath, "i", "", "Path to input JSON file")
	flag.StringVar(&inputPath, "input", "", "Path to input JSON file")
	flag.StringVar(&outputPath, "o", "", "Path to output JSON file")
	flag.StringVar(&outputPath, "output", "", "Path to output JSON file")
	flag.Parse()

	if help {
		fmt.Println("mockdata -i input.json -o output.json")
		os.Exit(0)
	}

	if inputPath == "" {
		fmt.Println("Input path is required")
		os.Exit(1)
	}

	if outputPath == "" {
		fmt.Println("Output path is required")
		os.Exit(1)
	}
}
