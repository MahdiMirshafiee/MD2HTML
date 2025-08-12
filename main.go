package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: md2html <your-file.md>")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	ext := strings.ToLower(filepath.Ext(inputFile))
	if ext == "md" {
		fmt.Fprintln(os.Stderr, "Error: input file must be a markdown file")
		os.Exit(1)
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: no such file or directory")
		os.Exit(1)
	}

	fmt.Println(string(content))
}
