# Stardew Valley Name Generator

This is a pretty direct port of Stardew Valley's random farm animal name generator, based on ~~c# code extracted from the game files~~ sheer intuition. Work is ongoing to clean up the generator code and ensure algorithm correctness.

## Command line tool

To install the command line tool, which just prints a name on a single line of standard output, run `go install` in the base of the directory.

## Library

To use as a software library, import `github.com/jemisonf/sdv_namegen/pkg/generator` and use the `GenerateName` function in your code. `sdv_namegen.go` is actually a pretty good example of what this might look like:

```go
package main

import (
	"fmt"

	"github.com/jemisonf/sdv_namegen/pkg/generator"
)

func main() {
	fmt.Println(generator.GenerateName())
}
```
