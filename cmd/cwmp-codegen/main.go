package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Niceblueman/cwmp-codegen/internal/generator"
	"github.com/Niceblueman/cwmp-codegen/internal/parser"
)

func main() {
	// Define command-line flags
	inputFile := flag.String("input", "", "Path to the XML model file (required)")
	outputLang := flag.String("lang", "golang", "Output language (golang, typescript, cheader)")
	outputDir := flag.String("output", "./output", "Directory for generated files")

	// Parse flags
	flag.Parse()

	// Validate required flags
	if *inputFile == "" {
		fmt.Println("Error: input file is required")
		flag.Usage()
		os.Exit(1)
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

	// Generate code based on the selected language
	fmt.Printf("Generating %s code...\n", *outputLang)
	var genErr error
	var outputFiles []string

	switch *outputLang {
	case "golang":
		outputFiles, genErr = generator.GenerateGolang(model, *outputDir)
	case "typescript":
		outputFiles, genErr = generator.GenerateTypeScript(model, *outputDir)
	case "cheader":
		outputFiles, genErr = generator.GenerateCHeader(model, *outputDir)
	default:
		fmt.Printf("Unsupported language: %s. Defaulting to golang.\n", *outputLang)
		outputFiles, genErr = generator.GenerateGolang(model, *outputDir)
	}

	if genErr != nil {
		fmt.Printf("Error generating code: %v\n", genErr)
		os.Exit(1)
	}

	// Report success
	fmt.Println("Code generation completed successfully!")
	fmt.Println("Generated files:")
	for _, file := range outputFiles {
		fmt.Println("-", filepath.Join(*outputDir, file))
	}
}
