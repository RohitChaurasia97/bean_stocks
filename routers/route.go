// MIT License

// Copyright (c) The RAI Authors

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package routers

import (
	"net/http"

	"bean_stocks/handlers"
	"bean_stocks/repositories"
	"bean_stocks/services"

	"github.com/labstack/echo/v4"
	"github.com/retail-ai-inc/bean"
)

type Repositories struct {
	exampleRepo                repositories.ExampleRepository
	stockstimelineddetailsRepo repositories.StockstimelineddetailsRepository // added by bean
	stockinfoRepo              repositories.StockinfoRepository              // added by bean
}

type Services struct {
	exampleSvc                services.ExampleService
	stockstimelineddetailsSvc services.StockstimelineddetailsService // added by bean
	stockinfoSvc              services.StockinfoService              // added by bean
}

type Handlers struct {
	exampleHdlr                handlers.ExampleHandler
	stockstimelineddetailsHdlr handlers.StockstimelineddetailsHandler // added by bean
	stockinfoHdlr              handlers.StockinfoHandler              // added by bean
}

func Init(b *bean.Bean) {

	e := b.Echo

	repos := &Repositories{
		exampleRepo:                repositories.NewExampleRepository(b.DBConn),
		stockstimelineddetailsRepo: repositories.NewStockstimelineddetailsRepository(b.DBConn), // added by bean
		stockinfoRepo:              repositories.NewStockinfoRepository(b.DBConn),              // added by bean
	}

	svcs := &Services{
		exampleSvc:                services.NewExampleService(repos.exampleRepo),
		stockstimelineddetailsSvc: services.NewStockstimelineddetailsService(repos.stockstimelineddetailsRepo), // added by bean
		stockinfoSvc:              services.NewStockinfoService(repos.stockinfoRepo),                           // added by bean
	}

	hdlrs := &Handlers{
		exampleHdlr:                handlers.NewExampleHandler(svcs.exampleSvc),
		stockstimelineddetailsHdlr: handlers.NewStockstimelineddetailsHandler(svcs.stockstimelineddetailsSvc), // added by bean
		stockinfoHdlr:              handlers.NewStockinfoHandler(svcs.stockinfoSvc),                           // added by bean
	}

	// Default index page goes to above JSON (/json) index page.
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": `bean_stocks ????`,
		})
	})

	// IMPORTANT: Just a JSON response index page. Please change or update it if you want.
	e.GET("/json", hdlrs.exampleHdlr.JSONIndex)

	// IMPORTANT: Just a HTML response index page. Please change or update it if you want.
	e.GET("/html", hdlrs.exampleHdlr.HTMLIndex)

	// Example of using validator.
	e.POST("/example", hdlrs.exampleHdlr.Validate)

	e.POST("/addstock", hdlrs.stockinfoHdlr.AddStockInfo)
	e.GET("/stocks", hdlrs.stockinfoHdlr.GetAllStockInfos)
	e.GET("/renderdata/:symbol", hdlrs.stockstimelineddetailsHdlr.RenderData)
	e.GET("/renderstock/:symbol", hdlrs.stockstimelineddetailsHdlr.RenderStock)

}
