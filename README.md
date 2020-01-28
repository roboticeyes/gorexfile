[![Godoc Reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/roboticeyes/gorexfile)
[![Build Status](https://travis-ci.org/roboticeyes/gorexfile.svg)](https://travis-ci.org/roboticeyes/gorexfile)
[![Go Report Card](https://goreportcard.com/badge/github.com/roboticeyes/gorexfile)](https://goreportcard.com/report/github.com/roboticeyes/gorexfile)

<p align="center">
  <img style="float: right;" src="assets/rex-go.png" alt="goREX logo"/>
</p>

# GoRexFile

The `gorexfile` library implements a Go reader and writer for the REXfile binary format. The REXfile
is used for [rexOS](https://www.rexos.org). The library can easily be integrated into your Go
project. It can help you to get started with REX as a developer.  For details about the binary file
format, please see the [format specification
document](https://github.com/roboticeyes/openrex/blob/master/doc/rex-spec-v1.md).

## Installation

> You can install Go by following [these instructions](https://golang.org/doc/install). Please note that Go >= 1.13. is required!

First, clone the repository to your local development path, and let go download all dependencies:

```
go mod tidy
```

This should download all required packages. To build all tools, you simple use the attached `Makefile` and call

```
make
```

## Usage

Make sure that you just include the `gorexfile` library in your application:

```go
package main

import (
    "github.com/roboticeyes/gorexfile/encoding/rexfile"
)
```

Please see the `examples` folder for further samples.

## Tools

### rxi

`rxi` is a simple command line tool which simply dumps REX file informations to the command line. It also allows to
extract images from the file directly. For more information, please call `rxi` directly.

# Todos

## REX File IO

* [ ] Data block text

## References

* [rexOS](https://www.rexos.org)
* [REX](https://app.rexos.cloud)
* [REX file format v1](https://github.com/roboticeyes/openrex/blob/master/doc/rex-spec-v1.md)
