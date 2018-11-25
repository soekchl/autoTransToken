package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/huobiapi/REST-GO-demo/untils"
	. "github.com/soekchl/myUtils"
	"golang.org/x/net/websocket"
)

func WebSocketAuth() {
	wsAuth, err := websocket.Dial(wsAuthUrl, "", origin)
	if err != nil {
		Error(err)
		return
	}
	defer func() {
		wsAuth.Close() //关闭连接
		wsAuth = nil
		WebSocketAuth()
	}()

	SendWsCmdAuth(getAuth()) // send auth
	//SendWsCmdAuth([]byte(`{"op":"sub", "topic":"accounts","cid":"1"}`)) // 账号资金变更 订阅事件
	SendWsCmdAuth([]byte(`{  "op":"sub", "topic": "orders.ethusdt", "cid":"2"}`)) // eth-usdt 交易变更 订阅事件(订单 提交和撤销)

	var msg = make([]byte, 65536)
	for {
		m, err := wsAuth.Read(msg)
		if err != nil {
			Error(err)
			return
		}
		tmBuff, err := decodeGzip(msg[:m])
		if err != nil {
			Error(err)
			continue
		}
		if m < 32 && strings.Index(string(tmBuff), "ping") >= 0 {
			SendWsCmdAuth([]byte(strings.Replace(string(tmBuff), "ping", "pong", 1)))
		} else {
			RecvWsAuthBuff <- tmBuff
		}
	}
}

func SendWsCmdAuth(cmd []byte) error {
	if wsAuth == nil {
		return WebSocketNull
	}
	_, err := wsAuth.Write(cmd)
	return err
}

func getAuth() []byte {
	mapParams2Sign := make(map[string]string)
	mapParams2Sign["AccessKeyId"] = access_key
	mapParams2Sign["SignatureMethod"] = "HmacSHA256"
	mapParams2Sign["SignatureVersion"] = "2"
	mapParams2Sign["Timestamp"] = time.Now().UTC().Format("2006-01-02T15:04:05")
	mapParams2Sign["Signature"] = untils.CreateSign(mapParams2Sign, "GET", "api.huobi.pro", "/ws/v1", secret_key)

	mapParams2Sign["op"] = "auth"
	buff, err := json.Marshal(mapParams2Sign)
	if err != nil {
		Error("getAuth ", err)
	}
	return buff
}
