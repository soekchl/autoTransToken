package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 下单
// placeRequestParams: 下单信息
// return: PlaceReturn对象
func Place(placeRequestParams PlaceRequestParams) PlaceReturn {
	placeReturn := PlaceReturn{}

	mapParams := make(map[string]string)
	mapParams["account-id"] = placeRequestParams.AccountID
	mapParams["amount"] = placeRequestParams.Amount
	if 0 < len(placeRequestParams.Price) {
		mapParams["price"] = placeRequestParams.Price
	}
	if 0 < len(placeRequestParams.Source) {
		mapParams["source"] = placeRequestParams.Source
	}
	mapParams["symbol"] = placeRequestParams.Symbol
	mapParams["type"] = placeRequestParams.Type

	strRequest := "/v1/order/orders/place"
	jsonPlaceReturn := ApiKeyPost(mapParams, strRequest)
	json.Unmarshal([]byte(jsonPlaceReturn), &placeReturn)

	return placeReturn
}

// 申请撤销一个订单请求
// strOrderID: 订单ID
// return: PlaceReturn对象
func SubmitCancel(strOrderID string) PlaceReturn {
	placeReturn := PlaceReturn{}

	strRequest := fmt.Sprintf("/v1/order/orders/%s/submitcancel", strOrderID)
	jsonPlaceReturn := ApiKeyPost(make(map[string]string), strRequest)
	json.Unmarshal([]byte(jsonPlaceReturn), &placeReturn)

	return placeReturn
}

func GetUsdtPrice() (float32, error) {
	response, err := http.Get(usdtPriceUrl)
	if err != nil {
		return 0, err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 {
		return 0, httpIsError
	}

	re := &ResultUSDT{}
	err = json.Unmarshal(body, re)
	if err != nil {
		return 0, err
	}

	if re.Code == 200 && len(re.Data) > 1 {
		return re.Data[0].Price, nil
	}
	return 0, nil
}
