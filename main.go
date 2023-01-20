//go:generate go run github.com/pointlander/peg example.peg

package main

import (
	"fmt"
	"os"
)

type Section struct {
	Name string
	Body string
}

type Document struct {
	Sections []*Section
}

func main() {
	text := `
[Foo]

content

[Bar]

content
`

	parser := &Parser{Buffer: text}
	parser.Init()
	err := parser.Parse()
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
	parser.Execute()

	for _, section := range parser.Document.Sections {
		fmt.Printf("\n=== %v ===\n", section.Name)
		fmt.Printf("%v", section.Body)
	}

}
