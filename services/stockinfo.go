package services

import (
	"context"

	// "github.com/retail-ai-inc/bean/trace"
	"bean_stocks/models"
	"bean_stocks/repositories"
)

type StockinfoService interface {
	AddStockinfoService(ctx context.Context, stockInfo models.StockInfo) (string, error)
	GetAllStockinfoService(ctx context.Context) ([]models.StockInfo, error)
}

type stockinfoService struct {
	stockinfoRepository repositories.StockinfoRepository
}

func NewStockinfoService(stockinfoRepo repositories.StockinfoRepository) *stockinfoService {
	return &stockinfoService{
		stockinfoRepository: stockinfoRepo,
	}
}

func (service *stockinfoService) AddStockinfoService(ctx context.Context, stockInfo models.StockInfo) (string, error) {
	return service.stockinfoRepository.AddStockInfoRepo(ctx, stockInfo)
}

func (service *stockinfoService) GetAllStockinfoService(ctx context.Context) ([]models.StockInfo, error) {
	return service.stockinfoRepository.GetAllStockInfoRepo(ctx)
}
