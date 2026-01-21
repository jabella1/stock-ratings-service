package restrequest

type GetListStockRatingRequest struct {
	Search         *string  `json:"search"`
	PageSize       *int32   `json:"pageSize"`
	PageNumber     *int32   `json:"pageNumber"`
	OrderBy        *string  `json:"orderBy"`
	OrderDirection *string  `json:"orderDirection"`
	MinUpside      *float32 `json:"minUpside"`
	MinPrice       *float64 `json:"minPrice"`
	MaxPrice       *float64 `json:"maxPrice"`
}
