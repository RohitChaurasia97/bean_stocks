package repositories

import (
	"bean_stocks/models"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// "github.com/retail-ai-inc/bean/trace"
	"github.com/retail-ai-inc/bean"
)

var apiKey = "f2841ae91c3f5db5fcccb4b5f1e1407f"

type StockinfoRepository interface {
	AddStockInfoRepo(c context.Context, stockInfo models.StockInfo) (string, error)
	GetAllStockInfoRepo(c context.Context) ([]models.StockInfo, error)
}

func NewStockinfoRepository(dbDeps *bean.DBDeps) *DbInfra {
	return &DbInfra{dbDeps}
}

func (db *DbInfra) AddStockInfoRepo(c context.Context, stockInfo models.StockInfo) (string, error) {
	// adding timelined stock data to db
	stockCode := stockInfo.Symbol
	fmpApi := fmt.Sprintf("https://fmpcloud.io/api/v3/historical-chart/1hour//%s?apikey=%s", stockCode, apiKey)
	// fmpApi := fmt.Sprintf("https://fmpcloud.io/api/v3/historical-chart/5min/%s?apikey=%s", stockCode, apiKey)
	resp, err := http.Get(fmpApi)
	if err != nil {
		return fmt.Sprintf("%s not added", stockCode), err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("%s not added", stockCode), err
	}

	var stockDetails []models.StockDetails
	json.Unmarshal(body, &stockDetails)
	db.Conn.MasterMySQLDB.Table(stockCode).AutoMigrate(&models.StockDetails{})

	for i, j := 0, len(stockDetails)-1; i < j; i, j = i+1, j-1 {
		stockDetails[i], stockDetails[j] = stockDetails[j], stockDetails[i]
	}

	for _, stockDetails := range stockDetails {
		var stockDetails = &models.StockDetails{Date: stockDetails.Date, Open: stockDetails.Open, Low: stockDetails.Low, High: stockDetails.High, Close: stockDetails.Close, Volume: stockDetails.Volume}
		err = db.Conn.MasterMySQLDB.Table(stockCode).Create(&stockDetails).Error
		if err != nil {
			return fmt.Sprintf("%s not added", stockCode), err
		}
	}

	// adding to list of stocks registered
	err = db.Conn.MasterMySQLDB.Create(&stockInfo).Error

	if err != nil {
		return fmt.Sprintf("%s not added", stockCode), err
	}
	return fmt.Sprintf("%s added successfully", stockCode), nil
}

func (db *DbInfra) GetAllStockInfoRepo(c context.Context) ([]models.StockInfo, error) {
	var stockInfos []models.StockInfo
	err := db.Conn.MasterMySQLDB.Find(&stockInfos).Error
	if err != nil {
		return []models.StockInfo{}, err
	}
	return stockInfos, nil
}
