package domain

type Operator string
type Direction string

const (
	Asc  Direction = "asc"
	Desc Direction = "desc"
)
const (
	Eq      Operator = "eq"      // Equals: field == value
	Ne      Operator = "ne"      // Not equals: field != value
	Gt      Operator = "gt"      // Greater than: field > value
	Lt      Operator = "lt"      // Less than: field < value
	Gte     Operator = "gte"     // Greater than or equal: field >= value
	Lte     Operator = "lte"     // Less than or equal: field <= value
	Like    Operator = "like"    // Pattern match: field LIKE value (e.g., SQL or partial match)
	In      Operator = "in"      // In list: field ∈ [v1, v2, ...]
	Nin     Operator = "nin"     // Not in list: field ∉ [v1, v2, ...]
	All     Operator = "all"     // Field (array) must contain all specified values
	Null    Operator = "null"    // Check for null: field IS NULL or IS NOT NULL
	Exists  Operator = "exists"  // Check if field exists (true/false)
	Regex   Operator = "regex"   // Regular expression match (useful in text fields or MongoDB)
	Between Operator = "between" // Value within range: [start, end] (e.g., date or price ranges)
)

// [field]:condition
type Filter map[string]Condition

// [operator]:values
type Condition map[Operator][]any

type Query struct {
	Filter     Filter
	Pagination *Pagination
}

type Sort struct {
	Key       string
	Direction Direction
}
type Pagination struct {
	Sorts  []Sort
	Offset uint
	Limit  uint
}

func NewQuery() *Query {
	return &Query{
		Filter: make(Filter),
		Pagination: &Pagination{
			Sorts:  make([]Sort, 0),
			Offset: 0,
			Limit:  0,
		},
	}
}
