package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestMainCLI(t *testing.T) {
	// Skip if short tests requested
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	// Build the command for testing
	tmpDir := t.TempDir()
	binPath := filepath.Join(tmpDir, "cwmp-codegen-test")

	buildCmd := exec.Command("go", "build", "-o", binPath)
	if output, err := buildCmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build test binary: %v\n%s", err, output)
	}

	// Create a test XML file
	testDataDir := filepath.Join(tmpDir, "testdata")
	if err := os.MkdirAll(testDataDir, 0755); err != nil {
		t.Fatalf("Failed to create test data directory: %v", err)
	}

	testFile := filepath.Join(testDataDir, "test.xml")
	xmlContent := `<?xml version="1.0" encoding="UTF-8"?>
<model name="TestIntegration" version="1.0">
  <description>Integration test model</description>
  <object name="IntegrationTest">
    <description>Test object for integration</description>
    <parameter name="TestParam" type="string">
      <description>Test parameter</description>
    </parameter>
  </object>
</model>`

	if err := os.WriteFile(testFile, []byte(xmlContent), 0644); err != nil {
		t.Fatalf("Failed to write test XML file: %v", err)
	}

	// Create output directories
	goOutDir := filepath.Join(tmpDir, "go-out")
	tsOutDir := filepath.Join(tmpDir, "ts-out")
	cOutDir := filepath.Join(tmpDir, "c-out")

	for _, dir := range []string{goOutDir, tsOutDir, cOutDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("Failed to create output directory %s: %v", dir, err)
		}
	}

	// Test Golang output
	goCmd := exec.Command(binPath, "--input", testFile, "--lang", "golang", "--output", goOutDir)
	if output, err := goCmd.CombinedOutput(); err != nil {
		t.Errorf("Golang generation failed: %v\n%s", err, output)
	}

	// Test TypeScript output
	tsCmd := exec.Command(binPath, "--input", testFile, "--lang", "typescript", "--output", tsOutDir)
	if output, err := tsCmd.CombinedOutput(); err != nil {
		t.Errorf("TypeScript generation failed: %v\n%s", err, output)
	}

	// Test C header output
	cCmd := exec.Command(binPath, "--input", testFile, "--lang", "cheader", "--output", cOutDir)
	if output, err := cCmd.CombinedOutput(); err != nil {
		t.Errorf("C header generation failed: %v\n%s", err, output)
	}

	// Verify output files exist
	goFile := filepath.Join(goOutDir, "testintegration.go")
	tsFile := filepath.Join(tsOutDir, "TestIntegration.ts")
	cFile := filepath.Join(cOutDir, "TestIntegration.h")

	for file, desc := range map[string]string{
		goFile: "Go",
		tsFile: "TypeScript",
		cFile:  "C header",
	} {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("%s output file was not generated", desc)
		}
	}
}
