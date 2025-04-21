# CWMP Data Model Converter

A CLI tool that converts CWMP (CPE WAN Management Protocol) data models from XML format to various programming languages.

## Features

- Converts CWMP XML models into:
  - Golang structs
  - TypeScript interfaces
  - C header files
- Easy-to-use CLI interface
- Preserves documentation and field types

## Installation

```bash
go install github.com/Niceblueman/cwmp-codegen/cmd/cwmp-codegen@latest