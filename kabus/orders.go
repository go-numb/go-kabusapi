package kabus

import "time"

// OrdersRequest - 注文約定照会のリクエストパラメータ
type OrdersRequest struct {
	Product Product // 取得する商品
}

// OrderResponse - 注文約定照会のレスポンス
type OrderResponse []Order

// Order - 注文約定照会で返される注文の情報
type Order struct {
	ID              string          `json:"ID"`              // 注文番号
	State           State           `json:"State"`           // 状態
	OrderState      OrderState      `json:"OrderState"`      // 注文状態
	OrdType         OrdType         `json:"OrdType"`         // 執行条件
	RecvTime        time.Time       `json:"RecvTime"`        // 受注日時
	Symbol          string          `json:"Symbol"`          // 銘柄コード
	SymbolName      string          `json:"SymbolName"`      // 銘柄名
	Exchange        Exchange        `json:"Exchange"`        // 市場コード
	ExchangeName    string          `json:"ExchangeName"`    // 市場名
	Price           float64         `json:"Price"`           // 値段
	OrderQty        float64         `json:"OrderQty"`        // 発注数量
	CumQty          float64         `json:"CumQty"`          // 約定数量
	Side            Side            `json:"Side"`            // 売買区分
	CashMargin      CashMargin      `json:"CashMargin"`      // 現物信用区分
	AccountType     AccountType     `json:"AccountType"`     // 口座種別
	DelivType       DelivType       `json:"DelivType"`       // 受渡区分
	ExpireDay       time.Time       `json:"ExpireDay"`       // 注文有効期限
	MarginTradeType MarginTradeType `json:"MarginTradeType"` // 信用取引区分
	Details         []OrderDetail   `json:"Details "`        // 注文詳細
}

// OrderDetail - 注文詳細
type OrderDetail struct {
	SeqNum        int       `json:"SeqNum"`        // 連番
	ID            int       `json:"ID"`            // 注文詳細番号
	RecType       RecType   `json:"RecType"`       // 明細種別
	ExchangeID    string    `json:"ExchangeID"`    // 取引所番号
	State         State     `json:"State"`         // 状態
	TransactTime  time.Time `json:"TransactTime"`  // 処理時刻
	OrdType       OrdType   `json:"OrdType"`       // 執行条件
	Price         float64   `json:"Price"`         // 値段
	Qty           float64   `json:"Qty"`           // 数量
	ExecutionID   string    `json:"ExecutionID"`   // 約定番号
	ExecutionDay  time.Time `json:"ExecutionDay"`  // 約定日時
	DelivDay      int       `json:"DelivDay"`      // 受渡日
	Commission    float64   `json:"Commission"`    // 手数料
	CommissionTax float64   `json:"CommissionTax"` // 手数料消費税
}
