# Go Calc

A small Go project focused on understanding core Go concepts through a simple price calculation workflow.

## Overview

This repository reads plain text price values from `prices.txt`, applies a set of tax rates, and writes the resulting tax-inclusive prices to JSON files.

The goal of the project is to practice and demonstrate:

- Go package structure
- interfaces and dependency injection
- pointer receivers vs value receivers
- file input/output
- JSON encoding
- error propagation
- data conversion and processing

## Project structure

- `main.go` - application entry point
- `prices/` - tax calculation logic and job struct
- `filemanager/` - file read/write implementation
- `iomanager/` - `IOManager` interface definition
- `conversion/` - utility to convert string lines to floats
- `prices.txt` - sample input file

## How it works

1. `main.go` defines several tax rates.
2. For each rate it creates a `FileManager` instance.
3. It passes that manager into `prices.NewTaxIncludePriceJob(...)`.
4. The job reads prices, converts them to floats, applies the tax, and writes JSON output.

## Run

From the project root:

```bash
go run main.go
```

This will generate output files like:

- `tax_include_prices_7.json`
- `tax_include_prices_8.json`
- `tax_include_prices_6.json`
- `tax_include_prices_5.json`

## Notes

- The project uses an interface-based design: `iomanager.IOManager` defines the contract for reading and writing.
- `filemanager.FileManager` implements that interface with pointer receivers.
- The sample input file must contain one price per line.

## Learning focus

The project is intentionally simple and designed to help you understand Go core principles rather than a complex business domain.
