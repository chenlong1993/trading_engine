package types

import (
	"github.com/shopspring/decimal"
	matching_types "github.com/yzimhao/trading_engine/v2/pkg/matching/types"
)

const (
	TOPIC_ORDER_NEW            = "order_new" // %s 表示交易对symbol
	TOPIC_NOTIFY_ORDER_CANCEL  = "notify_order_cancel"
	TOPIC_PROCESS_ORDER_CANCEL = "process_order_cancel"
	TOPIC_ORDER_TRADE          = "order_trade"
	TOPIC_ORDER_SETTLE         = "order_settle"
)

type EventOrderNew struct {
	Symbol    string                   `json:"symbol"`
	OrderId   string                   `json:"order_id"`
	OrderSide matching_types.OrderSide `json:"order_side"`
	OrderType matching_types.OrderType `json:"order_type"`
	Price     *decimal.Decimal         `json:"price"`
	Quantity  *decimal.Decimal         `json:"quantity"`
	Amount    *decimal.Decimal         `json:"amount"`
	MaxAmount *decimal.Decimal         `json:"max_amount"`
	MaxQty    *decimal.Decimal         `json:"max_qty"`
	NanoTime  int64                    `json:"nano_time"`
}

type EventOrderCancel struct{}

type EventOrderTrade struct {
	matching_types.TradeResult
}

type EventOrderSettle struct {
	matching_types.TradeResult
}
