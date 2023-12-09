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
SetBaseURL(string)
```

### SetTimeout

Set a timeout duration for all HTTP requests, the default is 30 seconds.

```go
SetTimeout(time.Duration)
```

## Usage

An example use of the library is below.

```go
package main

import (
	"flag"
	"log"

	"github.com/rydelll/papermc/client"
)

func main() {
	var version string
	var err error
	flag.StringVar(&version, "version", "latest", "version to download")

	c := client.NewClient()
	var err error

	if version == "latest" {
		version, err = c.Paper.LatestVersion()
		if err != nil {
			log.Fatal(err)
		}
	}

	info, err := c.Paper.LatestBuild(version)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Paper.Download(info)
	if err != nil {
		log.Fatal(err)
	}
}
```
