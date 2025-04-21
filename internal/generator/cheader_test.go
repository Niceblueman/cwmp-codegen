package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Niceblueman/cwmp-codegen/internal/models"
)

func TestGenerateCHeader(t *testing.T) {
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
	files, err := GenerateCHeader(model, tmpDir)
	if err != nil {
		t.Fatalf("GenerateCHeader returned error: %v", err)
	}

	// Check that we got the expected file
	if len(files) != 1 {
		t.Fatalf("Expected 1 file, got %d", len(files))
	}

	if files[0] != "TestModel.h" {
		t.Errorf("Expected filename 'TestModel.h', got '%s'", files[0])
	}

	// Read the generated file
	content, err := os.ReadFile(filepath.Join(tmpDir, files[0]))
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	// Verify the content
	contentStr := string(content)

	// Check for include guards
	if !strings.Contains(contentStr, "#ifndef TESTMODEL_H") {
		t.Error("Generated code doesn't contain expected include guard")
	}

	// Check for struct declaration
	if !strings.Contains(contentStr, "typedef struct TestObject") {
		t.Error("Generated code doesn't contain expected struct declaration")
	}

	// Check for field declarations
	if !strings.Contains(contentStr, "char* TestParam") {
		t.Error("Generated code doesn't contain expected string field")
	}

	if !strings.Contains(contentStr, "int32_t NumberParam") {
		t.Error("Generated code doesn't contain expected int field")
	}
}
