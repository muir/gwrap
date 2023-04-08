# list - golang generic wrappers for standard library functions

[![GoDoc](https://godoc.org/github.com/muir/gwrap?status.png)](https://pkg.go.dev/github.com/muir/gwrap)
![unit tests](https://github.com/muir/gwrap/actions/workflows/go.yml/badge.svg)
[![report card](https://goreportcard.com/badge/github.com/muir/gwrap)](https://goreportcard.com/report/github.com/muir/gwrap)
[![codecov](https://codecov.io/gh/muir/gwrap/branch/main/graph/badge.svg)](https://codecov.io/gh/muir/gwrap)

Install:

	go get github.com/muir/gwrap

---

This package is a collection of generic functions that wrap standard library functions.
Hopefully, this packge will quickly become obsolete because the library functions will
support generics directly.  Until then, there is this.

## SyncMap

SyncMap is a wrapper around sync.Map supporting the go 1.18 sync.Map

## CompareMap

CompareMap is a wrapper around sync.Map supporting the go 1.20 sync.Map. CompareMap is
only available when compiling with go 1.20 and above.
