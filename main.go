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
			background-color: #0d1117 !important;
			color: #c9d1d9 !important;
			font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
			max-width: 800px;
			margin: auto;
			padding: 2rem;
			line-height: 1.6;
		}
		h1, h2, h3 {
			color: #d1d5db !important;
			background-color: #161b22;
			padding: 10px;
			border-radius: 5px;
			margin-top: 1.5rem;
			margin-bottom: 1rem;
		}
		h1 {
			font-size: 2.5rem !important;
			text-align: center;
			border: 3px solid #9ca3af;
		}
		pre {
			background: #161b22 !important;
			padding: 1rem;
			overflow-x: auto;
			border: 2px solid #9ca3af;
			border-radius: 8px;
			color: #d1d5db;
		}
		code {
			font-family: monospace;
			background-color: #3a3a3b !important;
			color: #d1d5db !important;
			padding: 4px 6px;
			border-radius: 3px;
		}
		a {
			color: #d1d5db !important;
			text-decoration: none;
		}
		a:hover {
			text-decoration: underline;
		}
  		img {
    		max-width: 100%;
    		height: auto;
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
