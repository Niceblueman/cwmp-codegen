package models

import (
	"encoding/xml"
)

// DataModel represents the root element of a CWMP data model
type DataModel struct {
	XMLName     xml.Name    `xml:"model"`
	Name        string      `xml:"name,attr"`
	Description string      `xml:"description"`
	Version     string      `xml:"version,attr"`
	Objects     []Object    `xml:"object"`
	Parameters  []Parameter `xml:"parameter"`
}

// Object represents a CWMP object
type Object struct {
	Name          string      `xml:"name,attr"`
	Description   string      `xml:"description"`
	Access        string      `xml:"access,attr,omitempty"`
	MultiInstance bool        `xml:"multiInstance,attr,omitempty"`
	Objects       []Object    `xml:"object"`
	Parameters    []Parameter `xml:"parameter"`
}

// Parameter represents a CWMP parameter
type Parameter struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"description"`
	Type        string `xml:"type,attr"`
	Access      string `xml:"access,attr,omitempty"`
	Syntax      Syntax `xml:"syntax"`
}

// Syntax defines the value constraints for a parameter
type Syntax struct {
	Default string   `xml:"default,omitempty"`
	Enum    []string `xml:"enum>value,omitempty"`
	Range   []Range  `xml:"range,omitempty"`
	Pattern string   `xml:"pattern,omitempty"`
	Units   string   `xml:"units,attr,omitempty"`
	Size    Size     `xml:"size,omitempty"`
}

// Range defines min/max values for a parameter
type Range struct {
	Min string `xml:"min,attr"`
	Max string `xml:"max,attr"`
}

// Size defines size constraints for a parameter
type Size struct {
	Min int `xml:"min,attr"`
	Max int `xml:"max,attr,omitempty"`
}
