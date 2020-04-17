# Stardew Valley Name Generator [![Coverage Status](https://coveralls.io/repos/github/jemisonf/sdv_namegen/badge.svg?branch=master)](https://coveralls.io/github/jemisonf/sdv_namegen?branch=master)

This is a pretty direct port of Stardew Valley's random farm animal name generator, based on ~~c# code extracted from the game files~~ sheer intuition. Work is ongoing to clean up the generator code and ensure algorithm correctness.

## Command line tool

To install the command line tool, which just prints a name on a single line of standard output, run `go install` in the base of the directory.

## Library

To use as a software library, import `github.com/jemisonf/sdv_namegen/pkg/generator` and use the `GenerateName` function in your code. `sdv_namegen.go` provides a good example of what this might look like:

```go
package main

import (
	"fmt"

	"github.com/jemisonf/sdv_namegen/pkg/generator"
	"github.com/jemisonf/sdv_namegen/pkg/random"
)

func main() {
	random := random.NewRandom()
	fmt.Println(generator.GenerateName(random))
}

```

At this time, `GenerateName` must be called with a psuedorandom generator struct matching the `random.Generator` interface in `pkg/random`. This is used in this implementation to facilitate testing, as you can see in `pkg/generator/generator_test.go`. You may use the `random.NewRandom` function to create a generator initialized with the current time, or supply your own implementation.
