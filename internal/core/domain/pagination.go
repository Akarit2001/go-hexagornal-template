package domain

type Paginated[T any] struct {
	Items      []T
	TotalCount int64
}
