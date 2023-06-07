# gwrap - golang generic wrappers for standard library functions

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

## AtomicValue

AtomicValue is a wrapper around sync/atomic.Value.

## SyncPool

SyncPool is a wrapper for sync.Pool

## Heap

Heap is a heap-like wrapper for container/heap

## PriorityQueue

PriorityQueue is a wrapper for container/heap that implements priority queues.

PriorityQueue supports removing arbitrary items from the queue at any point. To
support that, the items in the queue must implement the `PQItem` interface. The
simplest way to implement that interface is to embed `PQItemEmbed` in the items
that will be in the priority queue.

