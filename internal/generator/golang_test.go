package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Niceblueman/cwmp-codegen/internal/models"
)

func TestGenerateGolang(t *testing.T) {
	// Create a test model
	model := &models.DataModel{
		Name:        "TestModel",
		Description: "Test model description",
		Version:     "1.0",
		Objects: []models.Object{
			{
				Name:        "TestObject",
				Description: "A test object",
				Parameters: []models.Parameter{
					{
						Name:        "TestParam",
						Description: "A test parameter",
						Type:        "string",
					},
					{
						Name:        "NumberParam",
						Description: "A number parameter",
						Type:        "int",
					},
				},
			},
		},
	}

	// Create temp output directory
	tmpDir := t.TempDir()

	// Generate the code
	files, err := GenerateGolang(model, tmpDir)
	if err != nil {
		t.Fatalf("GenerateGolang returned error: %v", err)
	}

	// Check that we got the expected file
	if len(files) != 1 {
		t.Fatalf("Expected 1 file, got %d", len(files))
	}

	if files[0] != "testmodel.go" {
		t.Errorf("Expected filename 'testmodel.go', got '%s'", files[0])
	}

	// Read the generated file
	content, err := os.ReadFile(filepath.Join(tmpDir, files[0]))
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	// Verify the content
	contentStr := string(content)

	// Check for package declaration
	if !strings.Contains(contentStr, "package testmodel") {
		t.Error("Generated code doesn't contain expected package declaration")
	}

	// Check for struct declaration
	if !strings.Contains(contentStr, "type TestObject struct {") {
		t.Error("Generated code doesn't contain expected struct declaration")
	}

	// Check for field declarations
	if !strings.Contains(contentStr, "TestParam string") {
		t.Error("Generated code doesn't contain expected string field")
	}

	if !strings.Contains(contentStr, "NumberParam int") {
		t.Error("Generated code doesn't contain expected int field")
	}
}

func TestConvertObjectToGoStruct(t *testing.T) {
	// Test with a simple object
	obj := models.Object{
		Name:        "SimpleObject",
		Description: "A simple test object",
		Parameters: []models.Parameter{
			{
				Name:        "StringParam",
				Description: "A string parameter",
				Type:        "string",
			},
			{
				Name:        "BoolParam",
				Description: "A boolean parameter",
				Type:        "boolean",
			},
		},
	}

	goStruct := convertObjectToGoStruct(obj)

	if goStruct.Name != "SimpleObject" {
		t.Errorf("Expected struct name 'SimpleObject', got '%s'", goStruct.Name)
	}

	if goStruct.GoName != "SimpleObject" {
		t.Errorf("Expected go struct name 'SimpleObject', got '%s'", goStruct.GoName)
	}

	if len(goStruct.Parameters) != 2 {
		t.Fatalf("Expected 2 parameters, got %d", len(goStruct.Parameters))
	}

	if goStruct.Parameters[0].GoType != "string" {
		t.Errorf("Expected parameter 0 type 'string', got '%s'", goStruct.Parameters[0].GoType)
	}

	if goStruct.Parameters[1].GoType != "bool" {
		t.Errorf("Expected parameter 1 type 'bool', got '%s'", goStruct.Parameters[1].GoType)
	}
}
