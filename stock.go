package main

import (
	"fmt"
	"strings"

	"github.com/imroc/req"
)

const POLL_URL string = "https://polling.finance.naver.com/api/realtime.nhn"

type StockInfo struct {
	name         string
	code         string
	currentValue int
	highestValue int
	lowestValue  int
	startValue   int
	prevValue    int
	changeAmount float64
	changeRate   float64
}

func poll(codes []string) ([]StockInfo, error) {
	param := req.Param{
		"query": fmt.Sprintf("SERVICE_ITEM:%s", strings.Join(codes, ",")),
	}
	resp, err := req.Get(POLL_URL, param)
	if err != nil {
		return nil, err
	}

	var dat map[string]interface{}

	err = resp.ToJSON(&dat)
	if err != nil {
		return nil, err
	}

	result := dat["result"].(map[string]interface{})
	areas := result["areas"].([]interface{})
	area := areas[0].(map[string]interface{})
	datas := area["datas"].([]interface{})

	stockInfos := []StockInfo{}
	for _, d := range datas {
		data := d.(map[string]interface{})
		code := data["cd"].(string)
		name := data["nm"].(string)
		currentValue := int(data["nv"].(float64))
		highestValue := int(data["hv"].(float64))
		lowestValue := int(data["lv"].(float64))
		startValue := int(data["ov"].(float64))
		prevValue := int(data["pcv"].(float64))

		dir := data["mt"].(string)
		changeAmount := data["cv"].(float64)
		changeRate := data["cr"].(float64)

		// dir: 1(-) / 2(+)
		if dir == "1" {
			changeAmount = -changeAmount
			changeRate = -changeRate
		}

		sinfo := StockInfo{
			code:         code,
			name:         name,
			currentValue: currentValue,
			highestValue: highestValue,
			lowestValue:  lowestValue,
			startValue:   startValue,
			prevValue:    prevValue,
			changeAmount: changeAmount,
			changeRate:   changeRate,
		}

		stockInfos = append(stockInfos, sinfo)
	}

	return stockInfos, nil
}
