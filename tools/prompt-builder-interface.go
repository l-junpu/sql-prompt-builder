package tools

/*
	Select
	Where
	Like
	From
	GroupBy
	Limit
	Offset
	Generate

	// advanced
	Case
	When
	Union

	Create
	Update
	Insert
	Delete
	Drop

	// tbc
	InnerJoin
*/

type QueryBuilder interface {
	Select(Columns ...string) QueryBuilder                                                      // Can add a select DISTINCT
	From(Table string) QueryBuilder                                                             // From Table
	Join(FirstTable string, SecondTable string, FirstCol string, SecondCol string) QueryBuilder // Does the INNERJOIN
	Where(Condition string, Comparison interface{}) QueryBuilder                                // Allow for nesting of Where conditions, check if previous command == WHERE, then we use an AND
	Like(Pattern string) QueryBuilder                                                           // Allows for nesting also, does not need an AND
	And(Condition string, Comparison interface{}) QueryBuilder                                  // Chains an AND
	Or(Condition string, Comparison interface{}) QueryBuilder                                   // Chains an OR
	GroupBy(Columns ...string) QueryBuilder                                                     // Separate each one by ,
	Limit(Rows int) QueryBuilder                                                                // Limits the number of rows
	Offset(Rows int) QueryBuilder
	Generate() string
}
