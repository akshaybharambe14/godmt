# GoDMT

[![Build](https://img.shields.io/github/workflow/status/averageflow/godmt/Test)](#)
[![Go Report Card](https://goreportcard.com/badge/github.com/averageflow/godmt)](https://goreportcard.com/report/github.com/averageflow/godmt)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/averageflow/godmt)](https://pkg.go.dev/github.com/averageflow/godmt)
[![Maintainability](https://api.codeclimate.com/v1/badges/8ee5c4680a29aef11331/maintainability)](https://codeclimate.com/github/averageflow/godmt/maintainability)
[![codecov](https://codecov.io/gh/averageflow/godmt/branch/master/graph/badge.svg?token=F4HW4K40T6)](https://codecov.io/gh/averageflow/godmt)
[![License](https://img.shields.io/github/license/averageflow/godmt.svg)](https://github.com/averageflow/godmt/blob/master/LICENSE.md)

GoDMT, the one and only Go Data Model Translator. The goal of this project is to provide a tool that can parse Go files
that include `var`, `const`, `map`, `struct` and `type` into an abstract syntax tree, aka AST.

<p align="center">
  <img width="250" height="150" src="web/DMT.png">
</p>

That AST will then be transformed into data models for several programming languages. Currently GoDMT can perform
translations to:

- TypeScript
- Swift (using Decodable structs)
- JSON
- PHP

Some small adjustments may need to be made to integrate the output into a project, but this should already save you a
lot of time and hassle, and will help you stay in sync with the Go version of your data models, in other languages.
Comments will be carried over 😉.

Currently, the supported operating systems are all of UNIX family:

- Linux
- BSD
- macOS

## Talk is cheap, show code

Feel free to browse some examples that I am happy to provide here:

- [Complex Structures](examples/ComplexStructures.md)
- [Constants](examples/Constants.md)
- [Maps](examples/Maps.md)
- [Pointers](examples/Pointers.md)
- [Slices](examples/Slices.md)
- [Structs](examples/Structs.md)

## Usage

See the [CLI Usage wiki page](https://github.com/averageflow/godmt/wiki/CLI-usage) for more details on using the tool.

See the [Tags and Name Conversion section](https://github.com/averageflow/godmt/wiki/Tags-and-name-conversion) to understand more about how entities get parsed and converted to other languages.

## Building

To build this application as a binary simply navigate to `cmd/godmt` and run `go build`.

