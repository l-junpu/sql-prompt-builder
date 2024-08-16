package main

import (
	"fmt"
	"sql-prompt-builder/tools"
)

func main() {
	b := tools.CreatePromptBuilder()

	query := b.Select("Column1").
		From("Table1").
		Join("Table1", "Table2", "Table1Column1", "Table2Column1").
		Join("Table2", "Table3", "Table2Column1", "Table3Column1").
		Where("age > %v", 30).Or("name = '%s'", "Jun Pu").
		GroupBy("name").
		Limit(10).Offset(5).
		Generate()

	fmt.Println(query)
}
