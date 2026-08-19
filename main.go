package main

import (
	"fmt"
	"os"
)

// template_engine - Project template generator
func template_engine(path string) {
	fmt.Println("========================================")
	fmt.Println("  Template-Engine")
	fmt.Println("  Project template generator")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	template_engine(path)
}
