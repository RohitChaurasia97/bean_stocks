package services

import (
	"context"

	// "github.com/retail-ai-inc/bean/trace"
	"bean_stocks/models"
	"bean_stocks/repositories"

	"github.com/yudai/pp"
)

type StockinfoService interface {
	StockinfoServiceExampleFunc(ctx context.Context) (string, error)
	AddStockinfoService(ctx context.Context, stockInfo models.StockInfo) ([]models.StockInfo, error)
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

func (service *stockinfoService) StockinfoServiceExampleFunc(ctx context.Context) (string, error) {
	// IMPORTANT: If you wanna trace the performance of your handler function then uncomment following 3 lines
	// finish := trace.Start(ctx, "http.service")
	// defer finish()
	return "StockinfoService", nil
}

func (service *stockinfoService) AddStockinfoService(ctx context.Context, stockInfo models.StockInfo) ([]models.StockInfo, error) {
	pp.Println("AddStockinfoService triigered")
	return service.stockinfoRepository.AddStockInfoRepo(ctx, stockInfo)
}

func (service *stockinfoService) GetAllStockinfoService(ctx context.Context) ([]models.StockInfo, error) {
	return service.stockinfoRepository.GetAllStockInfoRepo(ctx)
}
