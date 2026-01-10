package pagination

type PaginationMetadata struct {
	PageNumber            int32  `json:"pageNumber"`
	TotalPages            int32  `json:"totalPages"`
	PageSize              *int32 `json:"pageSize,omitempty"`
	TotalRecords          int64  `json:"totalRecords"`
	RecordsReturnedInPage int32  `json:"recordsReturnedInPage"`
	CanGoBack             bool   `json:"canGoBack"`
	CanGoForward          bool   `json:"canGoForward"`
}

func CreatePaginationMetadata(pageNumber int32, pageSize int32, totalRecords int64,
	recordsReturned int32,
) *PaginationMetadata {
	totalPages := int32((totalRecords + int64(pageSize) - 1) / int64(pageSize))
	return &PaginationMetadata{
		PageNumber:            pageNumber,
		TotalPages:            totalPages,
		PageSize:              &pageSize,
		TotalRecords:          totalRecords,
		RecordsReturnedInPage: recordsReturned,
		CanGoBack:             pageNumber > 1,
		CanGoForward:          pageNumber < totalPages,
	}
}
