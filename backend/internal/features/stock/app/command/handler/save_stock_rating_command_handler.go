package handler

import (
	"context"

	"github.com/jabella1/stock-ratings-service/internal/features/common/utils"
	"github.com/jabella1/stock-ratings-service/internal/features/common/wrapper"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/command"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/interfaces"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/entities"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/repositories"
	unitofwork "github.com/jabella1/stock-ratings-service/internal/features/stock/domain/unitOfWork"
)

type SaveStockRatingCommandHandler struct {
	stockRatingRepository repositories.StockRatingRepository
	unitOfWork            unitofwork.UnitOfWork
}

func CreateSaveStockRatingCommandHandler(stockRatingRepository repositories.StockRatingRepository,
	unitOfWork unitofwork.UnitOfWork) interfaces.SaveStockRatingCommandHandler {
	return &SaveStockRatingCommandHandler{
		stockRatingRepository: stockRatingRepository,
		unitOfWork:            unitOfWork,
	}
}

func (s *SaveStockRatingCommandHandler) SaveStockRating(context context.Context, saveStockRatingCommand *command.SaveStockRatingCommand) (*wrapper.Response[command.SaveStockRatingResult], error) {
	stockRatingEntity, err := entities.CreateStockRating(saveStockRatingCommand.Ticker, saveStockRatingCommand.Company,
		saveStockRatingCommand.Brokerage, &saveStockRatingCommand.Action, &saveStockRatingCommand.RatingFrom, &saveStockRatingCommand.RatingTo,
		&saveStockRatingCommand.TargetFrom, &saveStockRatingCommand.TargetTo,
		float32(utils.CalculatePercentageChange(&saveStockRatingCommand.CurrentPrice, &saveStockRatingCommand.TargetTo)),
		utils.CalculatePercentageChange(&saveStockRatingCommand.TargetFrom, &saveStockRatingCommand.TargetTo), saveStockRatingCommand.CurrentPrice)
	if err != nil {
		return nil, err
	}

	transaction, err := s.unitOfWork.BeginTransaction(context)
	if err != nil {
		return nil, err
	}
	defer transaction.Rollback(context)

	result, err := s.stockRatingRepository.SaveStockRating(context, transaction, stockRatingEntity)

	if err != nil {
		return nil, err
	}

	if result.PriceChanged || result.UpsideChanged {
		var stockRatingHistoryEntity, err = entities.CreateStockRatingHistory(
			result.ID, result.OldPrice, stockRatingEntity.GetCurrentPrice(), &result.OldUpside, stockRatingEntity.GetUpside())
		if err != nil {
			return nil, err
		}
		err = s.stockRatingRepository.SaveStockRatingHistory(context, transaction, stockRatingHistoryEntity)
		if err != nil {
			return nil, err
		}
	}

	if err := transaction.Commit(context); err != nil {
		return nil, err
	}

	return &wrapper.Response[command.SaveStockRatingResult]{
		Data: command.SaveStockRatingResult{
			Ticker:       stockRatingEntity.GetTicker(),
			Company:      stockRatingEntity.GetCompany(),
			Action:       stockRatingEntity.GetAction(),
			RatingFrom:   stockRatingEntity.GetRatingFrom(),
			RatingTo:     stockRatingEntity.GetRatingTo(),
			TargetFrom:   stockRatingEntity.GetTargetFrom(),
			TargetTo:     stockRatingEntity.GetTargetTo(),
			Upside:       stockRatingEntity.GetUpside(),
			CurrentPrice: stockRatingEntity.GetCurrentPrice(),
			CreateAt:     stockRatingEntity.GetCreateAt(),
		},
	}, nil
}
