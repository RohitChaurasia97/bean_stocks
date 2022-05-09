package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	//"github.com/retail-ai-inc/bean/trace"
	"bean_stocks/services"
)

type StockstimelineddetailsHandler interface {
	RenderStock(c echo.Context) error
	RenderData(c echo.Context) error
}

type stockstimelineddetailsHandler struct {
	stockstimelineddetailsService services.StockstimelineddetailsService
}

func NewStockstimelineddetailsHandler(stockstimelineddetailsSvc services.StockstimelineddetailsService) *stockstimelineddetailsHandler {
	return &stockstimelineddetailsHandler{stockstimelineddetailsSvc}
}

func (handler *stockstimelineddetailsHandler) RenderStock(c echo.Context) error {
	symbol := c.Param("symbol")
	return c.Render(http.StatusOK, "base", echo.Map{
		"symbol":   symbol,
		"template": "templates/render",
	})
}

func (handler *stockstimelineddetailsHandler) RenderData(c echo.Context) error {
	ctx := context.Background()
	symbol := c.Param("symbol")
	output, err := handler.stockstimelineddetailsService.RenderDataService(ctx, symbol)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, output)
}
