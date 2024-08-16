package tools

import (
	"fmt"
	"strings"
)

/*
	Functionality to support:
	- Basic Functions
		- Query Builder: Allow users to build SQL queries by selecting tables, columns, and conditions.
		- SQL Syntax Highlighting: Highlight SQL syntax to make it easier to read and write.
		- Auto-Completion: Provide auto-completion suggestions for table and column names, as well as SQL keywords.
		- Query Execution: Execute the built query and display the results.
	- Query Building Functions
		- Table Selection: Allow users to select tables from a database or schema.
		- Column Selection: Enable users to select columns from the chosen tables.
		- Join Support: Support various join types (e.g., INNER JOIN, LEFT JOIN, RIGHT JOIN, FULL OUTER JOIN).
		- Filtering: Allow users to add conditions (e.g., WHERE, HAVING) to filter data.
		- Sorting: Enable users to sort data by one or more columns.
		- Grouping: Support grouping data by one or more columns.
		- Aggregation: Allow users to apply aggregate functions (e.g., SUM, COUNT, AVG) to columns.
	- Advanced Functions
		- Subqueries: Support subqueries in the WHERE, FROM, or SELECT clauses.
		- Common Table Expressions (CTEs): Allow users to define CTEs for complex queries.
		- Window Functions: Support window functions (e.g., ROW_NUMBER, RANK, LAG) for advanced data analysis.
		- Full-Text Search: Enable users to perform full-text searches on columns.
		- Regular Expression Support: Allow users to use regular expressions in conditions.
*/

type PromptBuilder struct {
	Prompt string
}

func (b *PromptBuilder) Select(Columns ...string) QueryBuilder {
	for i, col := range Columns {
		if i == 0 {
			b.Prompt = fmt.Sprintf("SELECT %s", col)
		} else {
			b.Prompt = fmt.Sprintf("%s, %s", b.Prompt, col)
		}
	}
	b.Prompt = fmt.Sprintf("%s\n", b.Prompt)

	return b
}

func (b *PromptBuilder) From(Table string) QueryBuilder {
	b.Prompt = fmt.Sprintf("%sFROM %s\n", b.Prompt, Table)
	return b
}

func (b *PromptBuilder) Join(FirstTable string, SecondTable string, FirstCol string, SecondCol string) QueryBuilder {
	b.Prompt = fmt.Sprintf("%sINNER JOIN %s ON %s.%s = %s.%s\n", b.Prompt, SecondTable, FirstTable, FirstCol, SecondTable, SecondCol)
	return b
}

func (b *PromptBuilder) Where(Condition string, Comparison ...interface{}) QueryBuilder {
	Condition = fmt.Sprintf(Condition, Comparison...)
	b.Prompt = fmt.Sprintf("%sWHERE %s\n", b.Prompt, Condition)
	return b
}

func (b *PromptBuilder) Like(Pattern string) QueryBuilder {
	b.Prompt = strings.TrimSpace(b.Prompt)
	b.Prompt = fmt.Sprintf("%s LIKE %s\n", b.Prompt, Pattern)
	return b
}

func (b *PromptBuilder) And(Condition string, Comparison interface{}) QueryBuilder {
	b.Prompt = strings.TrimSpace(b.Prompt)
	Condition = fmt.Sprintf(Condition, Comparison)
	b.Prompt = fmt.Sprintf("%s AND %s\n", b.Prompt, Condition)
	return b
}

func (b *PromptBuilder) Or(Condition string, Comparison interface{}) QueryBuilder {
	b.Prompt = strings.TrimSpace(b.Prompt)
	Condition = fmt.Sprintf(Condition, Comparison)
	b.Prompt = fmt.Sprintf("%s OR %s\n", b.Prompt, Condition)
	return b
}

func (b *PromptBuilder) GroupBy(Columns ...string) QueryBuilder {
	for i, col := range Columns {
		if i == 0 {
			b.Prompt = fmt.Sprintf("%sGROUP BY %s", b.Prompt, col)
		} else {
			b.Prompt = fmt.Sprintf("%s, %s", b.Prompt, col)
		}
	}
	b.Prompt = fmt.Sprintf("%s\n", b.Prompt)
	return b
}

func (b *PromptBuilder) Limit(Rows int) QueryBuilder {
	b.Prompt = fmt.Sprintf("%sLIMIT %v\n", b.Prompt, Rows)
	return b
}

func (b *PromptBuilder) Offset(Rows int) QueryBuilder {
	trimmed_prompt := strings.TrimSpace(b.Prompt)
	ss := strings.Split(trimmed_prompt, "\n")

	if strings.Contains(ss[len(ss)-1], "LIMIT") {
		b.Prompt = strings.TrimSpace(b.Prompt)
		b.Prompt = fmt.Sprintf("%s OFFSET %v\n", b.Prompt, Rows)
	} else {
		b.Prompt = fmt.Sprintf("%sOFFSET %v\n", b.Prompt, Rows)
	}
	return b
}

func (b *PromptBuilder) Generate() string {
	b.Prompt = strings.TrimSpace(b.Prompt)
	return b.Prompt + ";"
}
