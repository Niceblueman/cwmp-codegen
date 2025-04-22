package models

import (
	"encoding/xml"
)

// Document represents the top-level XML element in a CWMP data model file
type Document struct {
	XMLName      xml.Name     `xml:"document"`
	Xmlns        string       `xml:"xmlns,attr,omitempty"`
	Spec         string       `xml:"spec,attr,omitempty"`
	DataTypes    []DataType   `xml:"dataType"`
	Bibliography Bibliography `xml:"bibliography"`
	Models       []DataModel  `xml:"model"`
}

// Bibliography contains references used in the document
type Bibliography struct {
	References []Reference `xml:"reference"`
}

// Reference represents a bibliographic reference
type Reference struct {
	ID           string `xml:"id,attr"`
	Name         string `xml:"name"`
	Title        string `xml:"title,omitempty"`
	Organization string `xml:"organization,omitempty"`
	Category     string `xml:"category,omitempty"`
	Date         string `xml:"date,omitempty"`
	Hyperlink    string `xml:"hyperlink,omitempty"`
}

// DataType represents a custom data type definition
type DataType struct {
	Name        string      `xml:"name,attr"`
	Description string      `xml:"description,omitempty"`
	String      *StringType `xml:"string,omitempty"`
}

// StringType represents string data type constraints
type StringType struct {
	Size    *Size     `xml:"size,omitempty"`
	Pattern []Pattern `xml:"pattern,omitempty"`
}

// Pattern represents a validation pattern
type Pattern struct {
	Value   string `xml:"value,attr,omitempty"`
	Content string `xml:",chardata"`
}

// DataModel represents a CWMP data model
type DataModel struct {
	XMLName     xml.Name    `xml:"model"`
	Name        string      `xml:"name,attr"`
	Description string      `xml:"description,omitempty"`
	Version     string      `xml:"version,attr,omitempty"`
	Objects     []Object    `xml:"object"`
	Parameters  []Parameter `xml:"parameter"`
}

// Object represents a CWMP object
type Object struct {
	Name                string      `xml:"name,attr"`
	Description         string      `xml:"description,omitempty"`
	Access              string      `xml:"access,attr,omitempty"`
	MinEntries          string      `xml:"minEntries,attr,omitempty"`
	MaxEntries          string      `xml:"maxEntries,attr,omitempty"`
	NumEntriesParameter string      `xml:"numEntriesParameter,attr,omitempty"`
	EnableParameter     string      `xml:"enableParameter,attr,omitempty"`
	UniqueKeys          []UniqueKey `xml:"uniqueKey"`
	Objects             []Object    `xml:"object"`
	Parameters          []Parameter `xml:"parameter"`
	MultiInstance       bool        // Derived field for code generation
	Path                string      // Full path including parent object paths
	HasIndexPlaceholder bool        // Whether the path contains an {i} placeholder
	ParentPath          string      // Path to parent object
	BaseName            string      // Name without the trailing dot (if any)
}

// GetPath returns the full path to this object
func (o *Object) GetPath() string {
	if o.Path != "" {
		return o.Path
	}
	return o.Name
}

// IsMultiInstance returns true if this object can have multiple instances
func (o *Object) IsMultiInstance() bool {
	return o.MultiInstance || o.HasIndexPlaceholder
}

// UniqueKey represents a unique key constraint
type UniqueKey struct {
	Parameters []ParameterRef `xml:"parameter"`
}

// ParameterRef is a reference to a parameter
type ParameterRef struct {
	Ref string `xml:"ref,attr"`
}

// Parameter represents a CWMP parameter
type Parameter struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"description,omitempty"`
	Access      string `xml:"access,attr,omitempty"`
	Syntax      Syntax `xml:"syntax"`
	Type        string // Derived field for code generation
	ParentPath  string // Path to parent object
	FullPath    string // Complete path including parent
}

// GetFullPath returns the full path to this parameter including parent paths
func (p *Parameter) GetFullPath() string {
	if p.FullPath != "" {
		return p.FullPath
	}
	if p.ParentPath != "" {
		return p.ParentPath + p.Name
	}
	return p.Name
}

// Syntax defines the value constraints for a parameter
type Syntax struct {
	Hidden      string       `xml:"hidden,attr,omitempty"`
	Default     string       `xml:"default,omitempty"`
	List        *List        `xml:"list,omitempty"`
	String      *StringCons  `xml:"string,omitempty"`
	Boolean     *Boolean     `xml:"boolean,omitempty"`
	DateTime    *DateTime    `xml:"dateTime,omitempty"`
	UnsignedInt *UnsignedInt `xml:"unsignedInt,omitempty"`
	DataTypeRef *DataTypeRef `xml:"dataType,omitempty"`
}

// List defines a list parameter
type List struct {
	Size *Size `xml:"size,omitempty"`
}

// StringCons defines string constraints
type StringCons struct {
	Size        *Size         `xml:"size,omitempty"`
	Enumeration []Enumeration `xml:"enumeration,omitempty"`
}

// Enumeration defines an enum value
type Enumeration struct {
	Value    string `xml:"value,attr"`
	Optional string `xml:"optional,attr,omitempty"`
}

// Boolean represents a boolean parameter type
type Boolean struct{}

// DateTime represents a dateTime parameter type
type DateTime struct{}

// UnsignedInt represents an unsignedInt parameter type
type UnsignedInt struct {
	Range *Range `xml:"range,omitempty"`
}

// Range defines min/max values for a parameter
type Range struct {
	MinInclusive string `xml:"minInclusive,attr,omitempty"`
	MaxInclusive string `xml:"maxInclusive,attr,omitempty"`
	Min          string `xml:"min,attr,omitempty"`
	Max          string `xml:"max,attr,omitempty"`
}

// DataTypeRef references a custom data type
type DataTypeRef struct {
	Ref string `xml:"ref,attr"`
}

// Size defines size constraints for a parameter
type Size struct {
	Min int `xml:"min,attr,omitempty"`
	Max int `xml:"maxLength,attr,omitempty"`
}
