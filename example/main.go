package main

import (
	"fmt"

	"github.com/epes/econfig"
)

type cfg struct {
	Parameter   string
	Other       int
	ID          int
	CamelCase   int
	Under_Score int
	Nested      struct {
		Name string
		Age  int
	}
}

func main() {
	var c cfg
	if err := econfig.Populate(&c, "config/", econfig.EnvironmentDevelopment); err != nil {
		panic(err)
	}

	fmt.Println(c)
}
