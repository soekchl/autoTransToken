package models

import (
	"autoTransToken/src/config"
	"errors"

	"golang.org/x/net/websocket"
)

const (
	SubEthKey = 1
	SubBtcKey = 2
	SubEosKey = 4
)

// API KEY
var (
	access_key string = ""
	secret_key string = ""

	// API请求地址, 不要带最后的/
	MARKET_URL string = ""
	TRADE_URL  string = ""
	HOST_NAME  string = ""

	origin       = "http://8.8.8.8"
	wsUrl        = "wss://api.huobi.pro/ws"
	wsAuthUrl    = "wss://api.huobi.pro/ws/v1"
	usdtPriceUrl = "https://otc-api.huobi.co/v1/data/trade-market?country=37&currency=1&payMethod=0&currPage=1&coinId=2&tradeType=buy&blockType=general&online=1"

	WebSocketNull = errors.New("WebSocket Is Null")
	httpIsError   = errors.New("Http Return Code Is Error")

	ws             *websocket.Conn = nil
	RecvWsBuff                     = make(chan []byte, 512)
	wsAuth         *websocket.Conn = nil
	RecvWsAuthBuff                 = make(chan []byte, 512)

	AccountId = 0
	Cid       = 0
)

func init() {
	config.InitConfig("huobi")
	loadConfig()

	if len(access_key) < 1 || len(secret_key) < 1 {
		panic("ACCESS_KEY or SECRET_KEY is null")
	}
}

func loadConfig() {
	config.ReLoadConfig()
	access_key = config.GetConfigString("ACCESS_KEY")
	secret_key = config.GetConfigString("SECRET_KEY")
	MARKET_URL = config.GetConfigString("MARKET_URL")
	TRADE_URL = config.GetConfigString("TRADE_URL")
	HOST_NAME = config.GetConfigString("HOST_NAME")
	origin = config.GetConfigString("origin")
	wsUrl = config.GetConfigString("wsUrl")
	wsAuthUrl = config.GetConfigString("wsAuthUrl")
	usdtPriceUrl = config.GetConfigString("usdtPriceUrl")
}
