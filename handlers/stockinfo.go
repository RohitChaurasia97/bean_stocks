package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	//"github.com/retail-ai-inc/bean/trace"
	"bean_stocks/models"
	"bean_stocks/services"
)

type StockinfoHandler interface {
	GetAllStockInfos(c echo.Context) error
	AddStockInfo(c echo.Context) error
}

type stockinfoHandler struct {
	stockinfoService services.StockinfoService
}

func NewStockinfoHandler(stockinfoSvc services.StockinfoService) *stockinfoHandler {
	return &stockinfoHandler{stockinfoSvc}
}

func (handler *stockinfoHandler) GetAllStockInfos(c echo.Context) error {
	ctx := context.Background()
	output, err := handler.stockinfoService.GetAllStockinfoService(ctx)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, output)
}

func (handler *stockinfoHandler) AddStockInfo(c echo.Context) error {
	ctx := context.Background()
	var stockInfo models.StockInfo
	err := c.Bind(&stockInfo)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	output, err := handler.stockinfoService.AddStockinfoService(ctx, stockInfo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"response": output,
	})
}
