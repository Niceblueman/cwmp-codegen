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
```

To run the tests:
bash
```bash
go test ./...
```
To run benchmarks:
bash
```bash
go test -bench=. ./...
```
To build the project:
bash
```bash
make build
```
To create a release with cross-platform binaries:
bash
```bash
make release
```