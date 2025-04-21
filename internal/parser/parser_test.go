package parser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseXML(t *testing.T) {
	// Create a temporary test file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_model.xml")
	
	xmlContent := `<?xml version="1.0" encoding="UTF-8"?>
<model name="TestDevice" version="1.0">
  <description>A test device model</description>
  <object name="Device" access="readOnly">
    <description>Device information</description>
    <parameter name="Manufacturer" access="readOnly" type="string">
      <description>The manufacturer</description>
    </parameter>
  </object>
</model>`

	if err := os.WriteFile(testFile, []byte(xmlContent), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Parse the test file
	model, err := ParseXML(testFile)
	if err != nil {
		t.Fatalf("Failed to parse XML: %v", err)
	}

	// Verify the parsed content
	if model.Name != "TestDevice" {
		t.Errorf("Expected model name 'TestDevice', got '%s'", model.Name)
	}
	
	if model.Version != "1.0" {
		t.Errorf("Expected version '1.0', got '%s'", model.Version)
	}
	
	if len(model.Objects) != 1 {
		t.Fatalf("Expected 1 object, got %d", len(model.Objects))
	}
	
	obj := model.Objects[0]
	if obj.Name != "Device" {
		t.Errorf("Expected object name 'Device', got '%s'", obj.Name)
	}
	
	if len(obj.Parameters) != 1 {
		t.Fatalf("Expected 1 parameter, got %d", len(obj.Parameters))
	}
	
	param := obj.Parameters[0]
	if param.Name != "Manufacturer" {
		t.Errorf("Expected parameter name 'Manufacturer', got '%s'", param.Name)
	}
	
	if param.Type != "string" {
		t.Errorf("Expected parameter type 'string', got '%s'", param.Type)
	}
}

func TestParseXMLInvalid(t *testing.T) {
	// Test with non-existent file
	_, err := ParseXML("non_existent_file.xml")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
	
	// Test with invalid XML content
	tmpDir := t.TempDir()
	invalidFile := filepath.Join(tmpDir, "invalid.xml")
	
	if err := os.WriteFile(invalidFile, []byte("This is not XML"), 0644); err != nil {
		t.Fatalf("Failed to create invalid test file: %v", err)
	}
	
	_, err = ParseXML(invalidFile)
	if err == nil {
		t.Error("Expected error for invalid XML, got nil")
	}
}