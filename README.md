# PaperMC Golang API Client

A PaperMC API wrapper for Golang.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Installation

To use the CLI tool

```
go install github.com/rydelll/papermc@latest
```

To use as a library

```
go get github.com/rydelll/papermc/client
```

## Options

### SetBaseURL

Set the PaperMC endpoint, only version 2 of the PaperMC API is currently supported; use at your own risk.

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
	"log"

	"github.com/rydelll/papermc/client"
)

func main() {
	c := client.New(client.WithTimeout(time.Second * 10))
	
	ver, err := c.Paper.Version.GetLatest()
	if err != nil {
		log.Fatal(err)
	}

	info, err := c.Paper.Build.GetLatest(ver)
	if err != nil {
		log.Fatal(err)
	}

	err := c.Paper.JAR.Download(info.Version, info.Build, info.JAR)
	if err != nil {
		log.Fatal(err)
	}

	err := c.Paper.JAR.ValidateChecksum(info.Checksum)
	if err != nil {
		log.Fatal(err)
	}
}
```
