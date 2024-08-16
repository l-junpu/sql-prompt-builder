# sql-prompt-builder

SQL Prompt Builder was developed as a standalone tool to help familiarize myself with the common SQL syntax. It is designed to generate SQL queries for common database operations, aiming to save time and reduce errors when writing SQL code. The project provides a programmable interface for building SQL queries, allowing users to create queries using a simple and intuitive API.

The tool addresses the problem of developers spending a significant amount of time searching for the correct SQL syntax to perform a specific operation, leading to wasted time, increased risk of errors, and inefficiencies in development workflows. By providing a simple, programmable way to generate SQL queries, SQL Prompt Builder streamlines development workflows, reduces errors, and improves efficiency.

## Supported Features

Data Retrieval
- SELECT
- FROM

Data Filtering
- WHERE
- LIKE
- AND
- OR

Data Joining
- JOIN

Data Aggregation
- GROUP BY

Data Limiting
- LIMIT
- OFFSET

## How To Use

### Example

```go
// Import the library
import (
    "sql-prompt-builder/tools"
)

func main() {
    // Create a SQL Prompt Builder
    b := tools.CreatePromptBuilder()

    // Construct Your SQL Query Programmatically
    query := b.Select("Column1").                                   // Select the column(s)
        From("Table1").                                             // Specify the Table(s)
        Join("Table1", "Table2", "Table1Column1", "Table2Column1"). // Specify additional Table(s)
        Join("Table2", "Table3", "Table2Column1", "Table3Column1"). // Specify additional Table(s)
        Where("age > %v", 30).Or("name = '%s'", "Jun Pu").          // Specify conditions
        GroupBy("name").                                            // Group results
        Limit(10).Offset(5).                                        // Limit results
        Generate()                                                  // Generate a SQL Query as a string

    // Display SQL Query
    fmt.Println(query)
}
```

### Example Output

```
SELECT Column1
FROM Table1
INNER JOIN Table2 ON Table1.Table1Column1 = Table2.Table2Column1
INNER JOIN Table3 ON Table2.Table2Column1 = Table3.Table3Column1
WHERE age > 30 OR name = 'Jun Pu'
GROUP BY name
LIMIT 10 OFFSET 5;
```