# BloomSky API in Go

[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![Build Status](https://travis-ci.org/patrickalin/bloomsky-api-go.svg?branch=master)](https://travis-ci.org/patrickalin/bloomsky-api-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/patrickalin/bloomsky-api-go)](https://goreportcard.com/report/github.com/patrickalin/bloomsky-api-go)
[![Coverage Status](https://coveralls.io/repos/github/patrickalin/bloomsky-api-go/badge.svg)](https://coveralls.io/github/patrickalin/bloomsky-api-go)
[![GoDoc](http://godoc.org/github.com/patrickalin/bloomsky-api-go?status.svg)](http://godoc.org/github.com/patrickalin/bloomsky-api-go)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![codebeat badge](https://codebeat.co/badges/f5a781ee-a438-40b7-b372-435401912239)](https://codebeat.co/projects/github-com-patrickalin-bloomsky-api-go-master)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/2886/badge)](https://bestpractices.coreinfrastructure.org/projects/2886)

## Package BloomskyStructure

The BloomskyStructure package provides APIs for bloomsky device.

## Bloomsky API Go

If you want, I have a [runtime client](https://github.com/patrickalin/bloomsky-client-go) which uses this API.

In other case, execute :

    go get github.com/patrickalin/bloomsky-api-go

## Folders / Files

* command/ : command to help me (build, test, ...)
* command/bench.sh : performance tests
* command/assembly.sh : add test and mock file in binnary
* command/build.sh : build source
* command/err.sh : check if all errors catch
* command/pprof.sh : performance tests
* command/pprofRaw.sh : performance tests
* command/tag.sh : tag in Git
* command/test.sh : test code
* command/torch.sh : Flame Graph, test performance
* mock/ : mock to simulate one bloomsky device
* scripts/ : use in Makefile
* testCase/ : test files use by tests
* .gitignore : ignore to commit
* .travis.yml : Continuous Integration
* LICENSE : text with license
* Makefile : command to help me (build, test, ...)
* README.md : this file
* bloomskyStructure.go : the main code
* bloomskyStructure_test.go : the test of the main code
* example_test.go : litte example to understand how to use the API
* utils.go : some reusable functions (error, log, ...). Not specific of this project.
* utils_test.go : file to test utils.go

## How to use API

Example in /example

    cd example
    go build .
    ./example

## How to test code

make test

or

command/test.sh

or

go test .

## License

The code is licensed under the permissive Apache v2.0 licence. This means you can do what you like with the software, as long as you include the required notices. [Read this](https://tldrlegal.com/license/apache-license-2.0-(apache-2.0)) for a summary.
