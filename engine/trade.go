package engine

import "encoding/json"

type Trade struct {
	TakerOrderID string `json:"taker_order_id"`
	MakerOrderID string `json:"maker_order_id"`
	Amount       uint64 `json:"amount"`
	Price        uint64 `json:"price"`
}

func (trade *Trade) FromJson(msg []byte) error {
	return json.Unmarshal(msg, trade)
}

func (trade *Trade) ToJson() []byte {
	str, _ := json.Marshal(trade)
	return str
}
