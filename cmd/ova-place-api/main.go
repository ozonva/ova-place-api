package main

import (
	"fmt"

	"github.com/ozonva/ova-place-api/internal/config"
)

func main() {
	fmt.Println("I'm an ova-place-api")

	config.Load()
}
