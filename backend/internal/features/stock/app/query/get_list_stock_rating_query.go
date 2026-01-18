package query

import (
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/dto"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/pagination"
)

type GetListStockRatingQuery struct {
	Search         *string
	PageSize       *int32
	PageNumber     *int32
	OrderBy        *string
	OrderDirection *string
	MinUpside      float32
	MinPrice       float64
	MaxPrice       float64
}

type GetListStockRatingResult struct {
	Results  *[]dto.GetListStockRatingResult
	Metadata *pagination.PaginationMetadata
}
