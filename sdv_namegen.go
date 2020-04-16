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
