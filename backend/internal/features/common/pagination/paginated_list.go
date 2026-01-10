package pagination

type PaginatedList[T any] struct {
	Results      *[]T
	TotalRecords int64
}
