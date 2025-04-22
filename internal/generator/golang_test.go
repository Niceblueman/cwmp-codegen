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

	// Check that we got the expected files (common_types.go + one per message)
	expectedFileCount := 1 + len(model.Objects) // common_types.go + one file per object
	if len(files) != expectedFileCount {
		t.Fatalf("Expected %d files, got %d", expectedFileCount, len(files))
	}

	// Verify common_types.go was generated
	if !contains(files, "common_types.go") {
		t.Errorf("Expected file 'common_types.go' not found in generated files")
	}

	// Verify TestObject.go was generated
	messageFile := "TestObject.go"
	if !contains(files, messageFile) {
		t.Errorf("Expected file '%s' not found in generated files", messageFile)
	}

	// Read the generated message file
	content, err := os.ReadFile(filepath.Join(tmpDir, messageFile))
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	// Verify the content
	contentStr := string(content)

	// Check for package declaration
	if !strings.Contains(contentStr, "package messages") {
		t.Error("Generated code doesn't contain expected package declaration")
	}

	// Check for struct declaration
	if !strings.Contains(contentStr, "type TestObject struct {") {
		t.Error("Generated code doesn't contain expected struct declaration")
	}

	// Check for method declarations
	if !strings.Contains(contentStr, "func NewTestObject()") {
		t.Error("Generated code doesn't contain expected constructor")
	}

	if !strings.Contains(contentStr, "func (msg *TestObject) GetID()") {
		t.Error("Generated code doesn't contain expected GetID method")
	}

	if !strings.Contains(contentStr, "func (msg *TestObject) GetName()") {
		t.Error("Generated code doesn't contain expected GetName method")
	}

	if !strings.Contains(contentStr, "func (msg *TestObject) CreateXML()") {
		t.Error("Generated code doesn't contain expected CreateXML method")
	}

	if !strings.Contains(contentStr, "func (msg *TestObject) Parse(") {
		t.Error("Generated code doesn't contain expected Parse method")
	}

	// Check for field declarations
	if !strings.Contains(contentStr, "TestParam string") {
		t.Error("Generated code doesn't contain expected string field")
	}

	if !strings.Contains(contentStr, "NumberParam int") {
		t.Error("Generated code doesn't contain expected int field")
	}

	// Read common_types.go
	commonContent, err := os.ReadFile(filepath.Join(tmpDir, "common_types.go"))
	if err != nil {
		t.Fatalf("Failed to read common_types.go: %v", err)
	}

	commonStr := string(commonContent)
	if !strings.Contains(commonStr, "type Message interface") {
		t.Error("common_types.go doesn't contain Message interface")
	}

	if !strings.Contains(commonStr, "type Envelope struct") {
		t.Error("common_types.go doesn't contain Envelope struct")
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

// Helper function to check if a slice contains a string
func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
