package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func Test_NewPositionsRequester(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg1 string
		arg2 bool
		want *positionsRequester
	}{
		{name: "本番用URLが取れる",
			arg1: "token1", arg2: true,
			want: &positionsRequester{httpClient: httpClient{url: "http://localhost:18080/kabusapi/positions", token: "token1"}}},
		{name: "検証用URLが取れる",
			arg1: "token2", arg2: false,
			want: &positionsRequester{httpClient: httpClient{url: "http://localhost:18081/kabusapi/positions", token: "token2"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := NewPositionsRequester(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_positionsRequester_Exec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *PositionsResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   positionsBody200,
			want1: &PositionsResponse{
				{
					ExecutionID:     "20200715E02N04738464",
					AccountType:     AccountTypeSpecific,
					Symbol:          "8306",
					SymbolName:      "三菱ＵＦＪフィナンシャル・グループ",
					Exchange:        StockExchangeToushou,
					ExchangeName:    "東証１部",
					ExecutionDay:    NewYmdNUM(time.Date(2020, 7, 2, 0, 0, 0, 0, time.Local)),
					Price:           704,
					LeavesQty:       500,
					HoldQty:         0,
					Side:            SideSell,
					Expenses:        0,
					Commission:      1620,
					CommissionTax:   162,
					ExpireDay:       NewYmdNUM(time.Date(2020, 12, 29, 0, 0, 0, 0, time.Local)),
					MarginTradeType: MarginTradeTypeSystem,
					CurrentPrice:    414.5,
					Valuation:       207250,
					ProfitLoss:      144750,
					ProfitLossRate:  41.12215909090909,
				},
			},
			want2: nil,
		},
		{name: "異常レスポンスをパースして返せる",
			status: http.StatusBadRequest,
			body:   `{"Code": 4001001,"Message": "内部エラー"}`,
			want1:  nil,
			want2: ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Body:       `{"Code": 4001001,"Message": "内部エラー"}`,
				Code:       4001001,
				Message:    "内部エラー",
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			}))
			defer ts.Close()

			req := &positionsRequester{httpClient{url: ts.URL}}
			got1, got2 := req.Exec(PositionsRequest{Product: ProductAll})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

const positionsBody200 = `[
  {
    "ExecutionID": "20200715E02N04738464",
    "AccountType": 4,
    "Symbol": "8306",
    "SymbolName": "三菱ＵＦＪフィナンシャル・グループ",
    "Exchange": 1,
    "ExchangeName": "東証１部",
    "ExecutionDay": 20200702,
    "Price": 704,
    "LeavesQty": 500,
    "HoldQty": 0,
    "Side": "1",
    "Expenses": 0,
    "Commission": 1620,
    "CommissionTax": 162,
    "ExpireDay": 20201229,
    "MarginTradeType": 1,
    "CurrentPrice": 414.5,
    "Valuation": 207250,
    "ProfitLoss": 144750,
    "ProfitLossRate": 41.12215909090909
  }
]`
