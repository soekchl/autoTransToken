package models

import (
	"strings"

	. "github.com/soekchl/myUtils"
	"golang.org/x/net/websocket"
)

func WebSocket(key int) {
	var err error
	ws, err = websocket.Dial(wsUrl, "", origin)
	if err != nil {
		Error(err)
		return
	}
	defer func() {
		ws.Close() //关闭连接
		ws = nil
		WebSocket(key)
	}()

	go setKline(key)

	var msg = make([]byte, 65536)
	for {
		m, err := ws.Read(msg)
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
			RecvWsBuff <- tmBuff
		}
	}
}

func setKline(key int) {
	if key&SubEthKey > 0 {
		SendWsCmd([]byte(`{"sub":"market.ethusdt.kline.1min","id": "id1"}`))
	}
	if key&SubEosKey > 0 {
		SendWsCmd([]byte(`{"sub":"market.eosusdt.kline.1min","id": "id2"}`))
	}
	if key&SubBtcKey > 0 {
		SendWsCmd([]byte(`{"sub":"market.Btcusdt.kline.1min","id": "id3"}`))
	}
}

func SendWsCmd(cmd []byte) error {
	if ws == nil {
		return WebSocketNull
	}
	_, err := ws.Write(cmd)
	return err
}
