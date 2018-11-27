package control

import (
	"autoTransToken/src/models"
	"encoding/json"
	"fmt"
	"strings"

	. "github.com/soekchl/myUtils"
)

func CheckData() {
	go models.WebSocketAuth()

	buff := make([]byte, 8096)
	for {
		select {
		case buff = <-models.RecvWsBuff:
		case buff = <-models.RecvWsAuthBuff:
			if strings.Contains(string(buff), "user-id") { // auth ok
				models.SendWsCmdAuth([]byte(`{  "op":"sub", "topic": "orders.ethusdt", "cid":"2"}`)) // eth-usdt 交易变更 订阅事件(订单 提交和撤销)

				// 获取历史委托信息
				models.SendWsCmdAuth([]byte(`{
  "op": "req",
"account-id" : 2165667,
  "topic": "orders.list",
  "cid": "3",
  "symbol": "ethusdt",
  "states": "submitted,partial-filled"
}`))
			}
		}
		if strings.Index(string(buff), "notify") >= 0 {
			tt := &Notify{}
			err := json.Unmarshal(buff, tt)
			if err != nil {
				Error(err)
				Notice("source buff = ", string(buff))
				continue
			}
			if tt.Data.OrderState == "filled" {
				Notice(fmt.Sprintf("出售成功 价格=%7v 数量=%4v 手续费用=%6v 种类=%v",
					tt.Data.OrderPrice,
					tt.Data.OrderAmount,
					tt.Data.OrderFees,
					tt.Data.Symbol,
				))
				continue
			}
		}
		Notice(string(buff))
	}
}
