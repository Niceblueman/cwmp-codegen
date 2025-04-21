package parser

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	"github.com/Niceblueman/cwmp-codegen/internal/models"
)

// ParseXML reads an XML file and converts it to our internal model representation
func ParseXML(filePath string) (*models.DataModel, error) {
	// Read XML file
	xmlFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	// Read the file content
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	// Parse XML into our data model
	var dataModel models.DataModel
	err = xml.Unmarshal(xmlData, &dataModel)
	if err != nil {
		return nil, err
	}

	return &dataModel, nil
}
