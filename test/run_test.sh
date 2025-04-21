#!/bin/bash

# Build the CLI tool
cd ..
go build -o cwmp-codegen ./cmd/cwmp-codegen

# Create test directories
mkdir -p test/output/golang
mkdir -p test/output/typescript
mkdir -p test/output/cheader

# Run the tool for each language
./cwmp-codegen --input=test/testdata/sample.xml --lang=golang --output=test/output/golang
./cwmp-codegen --input=test/testdata/sample.xml --lang=typescript --output=test/output/typescript
./cwmp-codegen --input=test/testdata/sample.xml --lang=cheader --output=test/output/cheader

echo "Test completed. Check the output directories for generated files."