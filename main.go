package main

import (
	"fmt"

	"github.com/nordicdyno/go-deps-A1/pkg/sharedapi"
)

func main() {
	api := sharedapi.NewAPI()
	fmt.Printf("API instance created: %T\n", api)
}
