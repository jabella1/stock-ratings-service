package interfaces

import (
	"context"

	"github.com/jabella1/stock-ratings-service/internal/features/common/wrapper"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/command"
)

type SaveStockRatingCommandHandler interface {
	SaveStockRating(context context.Context, saveStockRatingCommand *command.SaveStockRatingCommand) (*wrapper.Response[command.SaveStockRatingResult], error)
}
