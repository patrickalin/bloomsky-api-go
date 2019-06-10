#!/bin/bash
gofmt -d bloomskyStructure.go 
go list -f '{{ .Name }}: {{ .Doc }}'
errcheck ./...