package services

import (
	"context"

	// "github.com/retail-ai-inc/bean/trace"

	"bean_stocks/repositories"
)

type StockstimelineddetailsService interface {
	RenderDataService(ctx context.Context, symbol string) ([][]float32, error)
}

type stockstimelineddetailsService struct {
	stockstimelineddetailsRepository repositories.StockstimelineddetailsRepository
}

func NewStockstimelineddetailsService(stockstimelineddetailsRepo repositories.StockstimelineddetailsRepository) *stockstimelineddetailsService {
	return &stockstimelineddetailsService{
		stockstimelineddetailsRepository: stockstimelineddetailsRepo,
	}
}

func (service *stockstimelineddetailsService) RenderDataService(ctx context.Context, symbol string) ([][]float32, error) {
	return service.stockstimelineddetailsRepository.RenderDataRepo(ctx, symbol)
}
