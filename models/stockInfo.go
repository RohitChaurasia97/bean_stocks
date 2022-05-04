package models

type StockInfo struct {
	ID     uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

func (StockInfo) TableName() string {
	return "stockinfo"
}
