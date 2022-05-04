package repositories

import (
	"bean_stocks/models"
	"context"

	// "github.com/retail-ai-inc/bean/trace"
	"github.com/retail-ai-inc/bean"
	"github.com/yudai/pp"
)

type StockinfoRepository interface {
	StockinfoExampleFunc(ctx context.Context) (string, error)
	AddStockInfoRepo(c context.Context, stockInfo models.StockInfo) ([]models.StockInfo, error)
	GetAllStockInfoRepo(c context.Context) ([]models.StockInfo, error)
}

func NewStockinfoRepository(dbDeps *bean.DBDeps) *DbInfra {
	return &DbInfra{dbDeps}
}

func (db *DbInfra) StockinfoExampleFunc(ctx context.Context) (string, error) {
	// IMPORTANT: If you wanna trace the performance of your handler function then uncomment following 3 lines
	// finish := trace.Start(ctx, "db")
	// defer finish()
	return "Stockinfo", nil
}

func (db *DbInfra) AddStockInfoRepo(c context.Context, stockInfo models.StockInfo) ([]models.StockInfo, error) {
	pp.Println("stockInfo.Name", stockInfo.Name)
	pp.Println("stockInfo.Symbol", stockInfo.Symbol)
	// autoMigrateError := db.Conn.MasterMySQLDB.Table(stockInfo.Symbol).AutoMigrate(&stockInfo).Error
	// if autoMigrateError != nil {
	// 	pp.Println("autoMigrateError")
	// }
	pp.Println("AddStockInfoRepo triggered 2")
	err := db.Conn.MasterMySQLDB.Create(&stockInfo).Error
	pp.Println("AddStockInfoRepo triggered 3")
	var stockInfos []models.StockInfo
	db.Conn.MasterMySQLDB.Find(&stockInfos)
	if err != nil {
		return stockInfos, err
	}
	return stockInfos, nil
}

func (db *DbInfra) GetAllStockInfoRepo(c context.Context) ([]models.StockInfo, error) {
	var stockInfos []models.StockInfo
	pp.Println("GetAllStockInfoRepo triggered")
	err := db.Conn.MasterMySQLDB.Find(&stockInfos).Error
	pp.Println(stockInfos)
	if err != nil {
		return []models.StockInfo{}, err
	}
	return stockInfos, nil
}
