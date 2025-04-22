package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Niceblueman/cwmp-codegen/internal/generator"
	"github.com/Niceblueman/cwmp-codegen/internal/parser"
)

func main() {
	// Define command-line flags
	inputFile := flag.String("input", "", "Path to the XML model file (required)")
	outputDir := flag.String("output", "./output", "Directory for generated files")

	// Parse flags
	flag.Parse()

	// Validate required flags
	if *inputFile == "" {
		fmt.Println("Error: input file is required")
		flag.Usage()
		os.Exit(1)
	}

	// Fix output directory name if it has .go suffix (tr069.go -> tr069)
	if strings.HasSuffix(*outputDir, ".go") {
		*outputDir = strings.TrimSuffix(*outputDir, ".go")
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Parse the XML file
	fmt.Println("Parsing XML model:", *inputFile)
	model, err := parser.ParseXML(*inputFile)
	if err != nil {
		fmt.Printf("Error parsing XML: %v\n", err)
		os.Exit(1)
	}

	// Generate Golang code
	fmt.Println("Generating Golang code...")
	outputFiles, err := generator.GenerateGolang(model, *outputDir)
	if err != nil {
		fmt.Printf("Error generating code: %v\n", err)
		os.Exit(1)
	}

	// Report success
	fmt.Println("Code generation completed successfully!")
	fmt.Println("Generated files:")
	for _, file := range outputFiles {
		fmt.Println("-", filepath.Join(*outputDir, file))
	}
}
