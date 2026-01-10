package query

import (
	"github.com/jabella1/stock-ratings-service/internal/features/common/pagination"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/dto"
)

type GetListStockRatingQuery struct {
	Search         *string
	PageSize       *int32
	PageNumber     *int32
	OrderBy        *string
	OrderDirection *string
}

type GetListStockRatingResult struct {
	Results  *[]dto.GetListStockRatingResult
	Metadata *pagination.PaginationMetadata
}
