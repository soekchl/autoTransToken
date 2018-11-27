package control

type Notify struct {
	Data struct {
		Symbol      string
		OrderAmount string `json:"order-amount"`
		OrderPrice  string `json:"order-price"`
		OrderState  string `json:"order-state"`
		OrderFees   string `json:"filled-fees"`
	}
}
