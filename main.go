package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday/v2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: md2html <your-file.md>")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	ext := strings.ToLower(filepath.Ext(inputFile))
	if ext != ".md" {
		fmt.Fprintln(os.Stderr, "Error: input file must be a markdown file")
		os.Exit(1)
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: no such file or directory")
		os.Exit(1)
	}

	output := blackfriday.Run(content)
	htmlContent := string(output)

	fullPage := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<title>MD2HTML Output</title>
	<style>
	body {
		font-family: Arial, sans-serif;
		max-width: 800px;
		margin: auto;
		padding: 2rem;
		background: #f9f
		color: #333;
	}
	h1, h2, h3 {
		color: #005f99;
	}
	pre {
		background: #eee;
		padding: 1rem;
		overflow-x: auto;
	}
	code {
		font-family: monospace;
	}
	</style>
</head>
<body>` + htmlContent + `</body></html>`

	baseName := strings.TrimSuffix(inputFile, filepath.Ext(inputFile))
	outputFile := baseName + ".html"
	err = os.WriteFile(outputFile, []byte(fullPage), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: The file could not be written")
		os.Exit(1)
	}

	fmt.Printf("Converted HTML saved as %s\n", outputFile)
}
