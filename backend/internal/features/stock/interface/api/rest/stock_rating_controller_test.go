package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jabella1/stock-ratings-service/internal/features/common/utils"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/dto"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/query"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/pagination"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/interface/api/rest/restrequest"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGetStockRatingByTickerQueryHandler struct {
	mock.Mock
}

func (m *MockGetStockRatingByTickerQueryHandler) GetStockRatingByTicker(context context.Context, getStockRatingByTickerQuery *query.GetStockRatingByTickerQuery) (*query.GetStockRatingByTickerResult, error) {
	args := m.Called(context, getStockRatingByTickerQuery)
	return args.Get(0).(*query.GetStockRatingByTickerResult), args.Error(1)
}

type MockGetListStockRatingQueryHandler struct {
	mock.Mock
}

func (m *MockGetListStockRatingQueryHandler) GetListStockRating(context context.Context, getListStockRatingQuery *query.GetListStockRatingQuery) (*query.GetListStockRatingResult, error) {
	args := m.Called(context, getListStockRatingQuery)
	return args.Get(0).(*query.GetListStockRatingResult), args.Error(1)
}

func TestStockRatingController(t *testing.T) {
	e := echo.New()

	getStockRatingByTickerQueryHandler := new(MockGetStockRatingByTickerQueryHandler)
	getListStockRatingQueryHandler := new(MockGetListStockRatingQueryHandler)
	controller := CreateStockRatingController(e, getStockRatingByTickerQueryHandler, getListStockRatingQueryHandler)

	getStockRatingByTickerResult := &query.GetStockRatingByTickerResult{
		Result: &dto.GetStockRatingByTickerResult{
			ID:         1,
			Ticker:     "JFAM",
			Company:    "JuanFCompany",
			Brokerage:  utils.PtrString(""),
			Action:     utils.PtrString("target lowered by"),
			RatingFrom: utils.PtrString("Buy"),
			RatingTo:   utils.PtrString("Buy"),
			TargetFrom: 3.0,
			TargetTo:   2.5,
		},
	}

	getStockRatingByTickerQueryHandler.On("GetStockRatingByTicker", mock.Anything, mock.Anything).Return(getStockRatingByTickerResult, nil)

	request := httptest.NewRequest(http.MethodGet, "/api/v1/get-stock-rating-by-ticker/JFAM", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(request, rec)
	c.SetPath("/api/v1/get-stock-rating-by-ticker/:ticker")
	c.SetParamNames("ticker")
	c.SetParamValues("JFAM")

	err := controller.GetStockRatingByTicker(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "JFAM")

	request = httptest.NewRequest(http.MethodGet, "/api/v1/get-stock-rating-by-ticker/", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(request, rec)
	c.SetPath("/api/v1/get-stock-rating-by-ticker/:ticker")
	c.SetParamNames("ticker")
	c.SetParamValues("")

	err = controller.GetStockRatingByTicker(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "ticker parameter is required")

	getListStockRatingResult := &query.GetListStockRatingResult{
		Results: &[]dto.GetListStockRatingResult{{
			Ticker:       "JFAM",
			Company:      "JuanFCompany",
			Brokerage:    utils.PtrString(""),
			Action:       utils.PtrString("target lowered by"),
			RatingFrom:   utils.PtrString("Buy"),
			RatingTo:     utils.PtrString("Buy"),
			TargetFrom:   3.0,
			TargetTo:     2.5,
			Upside:       3.5,
			CurrentPrice: 5.0,
		}},
		Metadata: pagination.CreatePaginationMetadata(1, 10, 5, 5),
	}

	getListStockRatingQueryHandler.On("GetListStockRating", mock.Anything, mock.Anything).Return(getListStockRatingResult, nil)

	pageSize := int32(10)
	pageNumber := int32(1)

	reqBody := restrequest.GetListStockRatingRequest{
		PageSize:   &pageSize,
		PageNumber: &pageNumber,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	request = httptest.NewRequest(http.MethodPost, "/api/v1/get-list-stock-rating", bytes.NewReader(body))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(request, rec)

	err = controller.GetListStockRating(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "JFAM")

	request = httptest.NewRequest(http.MethodPost, "/api/v1/get-list-stock-rating", bytes.NewReader([]byte("invalid json")))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(request, rec)

	err = controller.GetListStockRating(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid request body")

	getStockRatingByTickerQueryHandler.AssertExpectations(t)
	getListStockRatingQueryHandler.AssertExpectations(t)
}
