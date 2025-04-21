package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Niceblueman/cwmp-codegen/internal/models"
)

func TestGenerateTypeScript(t *testing.T) {
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
	files, err := GenerateTypeScript(model, tmpDir)
	if err != nil {
		t.Fatalf("GenerateTypeScript returned error: %v", err)
	}

	// Check that we got the expected file
	if len(files) != 1 {
		t.Fatalf("Expected 1 file, got %d", len(files))
	}

	if files[0] != "TestModel.ts" {
		t.Errorf("Expected filename 'TestModel.ts', got '%s'", files[0])
	}

	// Read the generated file
	content, err := os.ReadFile(filepath.Join(tmpDir, files[0]))
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	// Verify the content
	contentStr := string(content)

	// Check for interface declaration
	if !strings.Contains(contentStr, "export interface TestObject") {
		t.Error("Generated code doesn't contain expected interface declaration")
	}

	// Check for property declarations
	if !strings.Contains(contentStr, "TestParam?: string") {
		t.Error("Generated code doesn't contain expected string property")
	}

	if !strings.Contains(contentStr, "NumberParam?: number") {
		t.Error("Generated code doesn't contain expected number property")
	}
}
