package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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

	if err := validateInput(inputPath); err != nil {
		fmt.Printf("invalid input: %s \n", err)
		os.Exit(0)
	}

	if err := validateOutput(outputPath); err != nil {
		fmt.Printf("invalid output: %s \n", err)
		os.Exit(0)
	}
}

func printUsage() {
	fmt.Println("Usage: mockdata [-i || --input] <input file> [-o || --output] <output file>")
	fmt.Println("-i --input: File input berupa JSON sebagai template")
	fmt.Println("-o --output: File output berupa JSON sebagai hasil")
}

// cek apakah file input tersebut ada atau tidak
func validateInput(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	return nil
}

// cek apakah file output tersebut ada atau tidak
func validateOutput(path string) error {
	if _, err := os.Stat(path); os.IsExist(err) {
		return nil
	}

	fmt.Println("File sudah ada di lokasi")

	confirmOverwrite()
	return nil
}

func confirmOverwrite() {
	fmt.Println("Apakah anda ingin menimpa file yang sudah ada (y/t)")

	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.ToLower(strings.TrimSpace(response))

	if response != "y" && response != "yes" && response != "ya" {
		fmt.Println("Membatalkan proses...")
		os.Exit(0)
	}
}
