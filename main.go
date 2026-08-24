package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: template-engine <template> <key=val> [key=val...]")
		os.Exit(1)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	result := string(data)
	for _, kv := range os.Args[2:] {
		parts := strings.SplitN(kv, "=", 2)
		if len(parts) == 2 {
			result = strings.ReplaceAll(result, "{{"+parts[0]+"}}", parts[1])
			result = strings.ReplaceAll(result, "{{ "+parts[0]+" }}", parts[1])
		}
	}
	remaining := strings.Count(result, "{{")
	fmt.Print(result)
	fmt.Fprintf(os.Stderr, "rendered (%d unresolved)\n", remaining)
}
