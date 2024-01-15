# PaperMC Golang API Client

A PaperMC API wrapper for Golang.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Installation

To use the CLI tool

```
go install github.com/rydelll/papermc
```

To use as a library

```
go get github.com/rydelll/papermc/client
```

## Options

### SetBaseURL

Set the PaperMC endpoint, only version 2 of the PaperMC API is currently supported so use at your own risk.

```go
WithBaseURL(string)
```

### SetTimeout

Set a timeout duration for all HTTP requests, the default is 30 seconds.

```go
WithTimeout(time.Duration)
```

## Usage

An example use of the library is given below.

```go
package main

import (
	"flag"
	"log"

	"github.com/rydelll/papermc/client"
)

func main() {
	
}
```
