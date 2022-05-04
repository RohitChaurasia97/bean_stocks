package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/retail-ai-inc/bean/async"
	berror "github.com/retail-ai-inc/bean/error"

	//"github.com/retail-ai-inc/bean/trace"
	"bean_stocks/models"
	"bean_stocks/services"
)

type StockinfoHandler interface {
	StockinfoJSONResponse(c echo.Context) error     // An example JSON response handler function
	StockinfoHTMLResponse(c echo.Context) error     // An example HTML response handler function
	StockinfoValidateResponse(c echo.Context) error // An example response handler function with validation
	AddStockInfo(c echo.Context) error
	GetAllStockInfos(c echo.Context) error
}

type stockinfoHandler struct {
	stockinfoService services.StockinfoService
}

func NewStockinfoHandler(stockinfoSvc services.StockinfoService) *stockinfoHandler {
	return &stockinfoHandler{stockinfoSvc}
}

func (handler *stockinfoHandler) StockinfoJSONResponse(c echo.Context) error {

	// IMPORTANT: If you wanna trace the performance of your handler function then uncomment following 3 lines
	// tctx := trace.NewTraceableContext(c.Request().Context())
	// finish := trace.Start(tctx, "http.handler")
	// defer finish()

	output, err := handler.stockinfoService.StockinfoServiceExampleFunc(c.Request().Context())
	if err != nil {
		return err
	}

	// IMPORTANT: Panic inside a goroutine will crash the whole application.
	// Example: How to execute some asynchronous code safely instead of plain goroutine:
	async.Execute(func(c echo.Context) {
		c.Logger().Debug(output)
		// IMPORTANT: Using sentry directly in goroutine may cause data race!
		// Need to create a new hub by cloning the existing one.
		// Example: How to use sentry safely in goroutine.
		// localHub := sentry.CurrentHub().Clone()
		// localHub.CaptureMessage(output)
	}, c.Echo())

	return c.JSON(http.StatusOK, map[string]string{
		"output": output,
	})
}

func (handler *stockinfoHandler) StockinfoHTMLResponse(c echo.Context) error {
	return c.Render(http.StatusOK, "index", echo.Map{
		"title": "Index title!",
		"add": func(a int, b int) int {
			return a + b
		},
		"test": map[string]interface{}{
			"a": "hi",
			"b": 10,
		},
		"copyrightYear": time.Now().Year(),
		"template":      "templates/master",
	})
}

func (handler *stockinfoHandler) StockinfoValidateResponse(c echo.Context) error {
	var params struct {
		Example string `json:"example" validate:"required,example,min=7"`
		Other   int    `json:"other" validate:"required,gt=0"`
	}

	if err := c.Bind(&params); err != nil {
		return berror.NewAPIError(http.StatusBadRequest, berror.PROBLEM_PARSING_JSON, err)
	}

	if err := c.Validate(params); err != nil {
		return err
	}

	return c.String(http.StatusOK, params.Example+" OK\n")
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
	return c.JSON(http.StatusOK, output)
}
