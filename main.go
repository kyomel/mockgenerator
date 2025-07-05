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

	if help || inputPath == "" || outputPath == "" {
		printUsage()
		os.Exit(0)
	}

}

func printUsage() {
	fmt.Println("Usage: mockdata [-i || --input] <input file> [-o || --output] <output file>")
	fmt.Println("-i --input: File input berupa JSON sebagai template")
	fmt.Println("-o --output: File output berupa JSON sebagai hasil")
}
