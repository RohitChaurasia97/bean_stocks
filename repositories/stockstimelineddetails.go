package repositories

import (
	"bean_stocks/models"
	"context"

	// "github.com/retail-ai-inc/bean/trace"
	"github.com/araddon/dateparse"
	"github.com/iancoleman/orderedmap"
	"github.com/retail-ai-inc/bean"
)

type StockstimelineddetailsRepository interface {
	RenderDataRepo(c context.Context, stockCode string) ([][]float32, error)
}

func NewStockstimelineddetailsRepository(dbDeps *bean.DBDeps) *DbInfra {
	return &DbInfra{dbDeps}
}

func getClosingDataForEachDate(stockTimelinedDetails []models.StockDetails) [][]float32 {
	// gets closing price for each date to render it in highcharts graph
	datesClosePriceMap := orderedmap.New()

	for _, details := range stockTimelinedDetails {
		date := details.Date[:10]
		datesClosePriceMap.Set(date, details.Close)
	}

	detailsToRender := make([][]float32, 2)
	keys := datesClosePriceMap.Keys()
	for _, dateStr := range keys {
		closePrice, _ := datesClosePriceMap.Get(dateStr)
		dateTime, _ := dateparse.ParseLocal(dateStr)
		timestamp := dateTime.Unix()
		detailsToRender = append(detailsToRender, []float32{float32(timestamp) * 1000, closePrice.(float32)})
	}
	detailsToRender = detailsToRender[2:]
	return detailsToRender
}

func (db *DbInfra) RenderDataRepo(ctx context.Context, stockCode string) ([][]float32, error) {
	var stockTimelinedDetails []models.StockDetails
	err := db.Conn.MasterMySQLDB.Table(stockCode).Find(&stockTimelinedDetails).Error

	detailsToRender := getClosingDataForEachDate(stockTimelinedDetails)
	getClosingDataForEachDate(stockTimelinedDetails)
	if err != nil {
		return detailsToRender, err
	}
	return detailsToRender, nil
}
