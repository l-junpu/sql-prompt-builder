package main

import (
	"fmt"
	"sql-Prompt-builder/tools"
)

func main() {
	b := tools.PromptBuilder{
		Prompt: "",
	}

	query := b.Select("Column1").
		From("Table1").
		Join("Table1", "Table2", "Table1Column1", "Table2Column1").
		Join("Table2", "Table3", "Table2Column1", "Table3Column1").
		Generate()

	fmt.Println(query)
}
