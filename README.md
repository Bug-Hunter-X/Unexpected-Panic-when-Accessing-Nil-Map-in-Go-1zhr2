# Unexpected Panic when Accessing Nil Map in Go

This repository demonstrates a common error in Go: accessing a key in a nil map.  Attempting this directly will result in a runtime panic.  This example shows the problem and a solution using proper initialization.

## Bug

The `bug.go` file contains code that demonstrates the problem:  the map `m` is not initialized, so accessing `m["key"]` causes a panic.

## Solution

The `bugSolution.go` file shows how to correctly initialize the map before use, preventing the panic.  This involves creating the map with `make(map[string]int)` or using a map literal.

This is a crucial concept for writing robust Go code.  Always ensure maps are initialized before use to avoid runtime panics.