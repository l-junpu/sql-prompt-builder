package builder

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

type PromptBuilderInterface interface {
	DataRetrieval
	DataFilters
	DataGrouping
	DataAggregation
	DataLimiters
	DataExtractor
}

type PromptBuilder struct {
	Prompt string
}

func CreatePromptBuilder() *PromptBuilder {
	b := PromptBuilder{
		Prompt: "",
	}
	return &b
}

/*
	-> Filtering & Sorting
	OrderBy
	Not
	In
	Between
	Like
	Not In
	Not Between
	Is Null / No Null

	-> Aggregation & Grouping
	Groupby
	Sum
	Avg
	Max
	Min
	Count
	Having

	-> Joining
	Inner Join (D)
	Left Join
	Right Join
	Full Outer Join
	Subquery
*/
