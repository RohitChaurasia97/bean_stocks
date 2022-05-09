package models

type StockDetails struct {
	Id     uint64  `gorm:"primary_key;AUTO_INCREMENT"`
	Date   string  `json:"date"`
	Open   float32 `json:"open"`
	Low    float32 `json:"low"`
	High   float32 `json:"high"`
	Close  float32 `json:"close"`
	Volume float32 `json:"volune"`
}
