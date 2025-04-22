package parser

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/Niceblueman/cwmp-codegen/internal/models"
)

// ParseXML reads an XML file or URL and converts it to our internal model representation
func ParseXML(source string) (*models.DataModel, error) {
	var xmlData []byte
	var err error

	// Check if the source is a URL or a local file
	if isURL(source) {
		xmlData, err = fetchFromURL(source)
	} else {
		xmlData, err = readFromFile(source)
	}

	if err != nil {
		return nil, err
	}

	// Parse XML into our document structure
	var document models.Document
	err = xml.Unmarshal(xmlData, &document)
	if err != nil {
		return nil, err
	}

	// Check if we have any models
	if len(document.Models) == 0 {
		return nil, errors.New("no models found in the document")
	}

	// Process the model to set derived fields
	processModel(&document.Models[0])

	// Return the first model found (most documents only have one)
	return &document.Models[0], nil
}

// processModel processes a data model to set derived fields
func processModel(model *models.DataModel) {
	// First pass: process all objects to set basic fields
	for i := range model.Objects {
		processObjectInitial(&model.Objects[i], "")
	}

	// Second pass: process all parameters to ensure access to full paths
	for i := range model.Objects {
		processObjectParameters(&model.Objects[i])
	}

	// Process any top-level parameters
	for i := range model.Parameters {
		processParameter(&model.Parameters[i])
	}
}

// processObjectInitial sets up the object hierarchy and basic fields on first pass
func processObjectInitial(obj *models.Object, parentPath string) {
	// Store the base name (without trailing dot)
	if strings.HasSuffix(obj.Name, ".") {
		obj.BaseName = obj.Name[:len(obj.Name)-1]
	} else {
		obj.BaseName = obj.Name
	}

	// Set the parent path
	obj.ParentPath = parentPath

	// Build the full path for this object
	obj.Path = parentPath + obj.Name

	// Check if this is a multi-instance object
	obj.HasIndexPlaceholder = strings.Contains(obj.Name, "{i}")
	obj.MultiInstance = obj.MaxEntries == "unbounded" || obj.MaxEntries != "1" || obj.HasIndexPlaceholder

	// Process nested objects recursively
	for i := range obj.Objects {
		processObjectInitial(&obj.Objects[i], obj.Path)
	}
}

// processObjectParameters processes parameters for an object after paths are set up
func processObjectParameters(obj *models.Object) {
	// Set path info for parameters
	for i := range obj.Parameters {
		obj.Parameters[i].ParentPath = obj.Path
		obj.Parameters[i].FullPath = obj.Path + obj.Parameters[i].Name
		processParameter(&obj.Parameters[i])
	}

	// Process nested objects recursively
	for i := range obj.Objects {
		processObjectParameters(&obj.Objects[i])
	}
}

// processObject processes an object to set derived fields and processes its children
// Kept for backward compatibility
func processObject(obj *models.Object) {
	// Determine if the object is multi-instance based on minEntries and maxEntries
	// If maxEntries is > 1 or unbounded, it's multi-instance
	obj.MultiInstance = obj.MaxEntries == "unbounded" || obj.MaxEntries != "1"

	// Check for {i} placeholder in name
	obj.HasIndexPlaceholder = strings.Contains(obj.Name, "{i}")

	// Process nested objects and parameters
	for i := range obj.Objects {
		processObject(&obj.Objects[i])
	}
	for i := range obj.Parameters {
		processParameter(&obj.Parameters[i])
	}
}

// processParameter processes a parameter to set derived fields
func processParameter(param *models.Parameter) {
	// Determine parameter type based on syntax
	syntax := param.Syntax

	if syntax.Boolean != nil {
		param.Type = "boolean"
	} else if syntax.String != nil {
		param.Type = "string"
	} else if syntax.DateTime != nil {
		param.Type = "datetime"
	} else if syntax.UnsignedInt != nil {
		param.Type = "unsignedInt"
	} else if syntax.DataTypeRef != nil {
		param.Type = syntax.DataTypeRef.Ref
	} else if syntax.List != nil {
		param.Type = "list"
	} else {
		param.Type = "string" // Default type
	}
}

// isURL checks if the input string is a URL
func isURL(input string) bool {
	input = strings.TrimSpace(input)
	u, err := url.Parse(input)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// fetchFromURL downloads content from a URL
func fetchFromURL(urlStr string) ([]byte, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP request failed with status: " + resp.Status)
	}

	return io.ReadAll(resp.Body)
}

// readFromFile reads content from a local file
func readFromFile(filePath string) ([]byte, error) {
	// Read XML file
	xmlFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	// Read the file content
	return io.ReadAll(xmlFile)
}
