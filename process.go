package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type streamData struct {
	Estr         string      `json:"e"` // 事件类型
	E            int         `json:"E"` // 事件时间
	CTime        int         `json:"C"` // 事件时间
	PriceChanged string      `json:"p"`
	Pairs        string      `json:"s"` // 交易对 BNBBTC  BTCUSDT
	Current      string      `json:"c"` // 最新成交价格
	Open         string      `json:"o"` // 24小时前开始第一笔成交价格
	OTime        interface{} `json:"O"` // 统计开始时间
	//High    string      `json:"h"` // 24小时内最高成交价
	//Low     interface{} `json:"l"` // 24小时内最低成交价
	//Volumn  string      `json:"v"` // 成交量
	Quality string `json:"q"` // 成交额
	Change  string `json:"P"`
}

func (s streamData) GetQuality() float64 {
	v, _ := strconv.ParseFloat(s.Quality, 64)
	return v
}

func (s streamData) GetCurrent() float64 {
	v, _ := strconv.ParseFloat(s.Current, 64)
	return v
}

func (s streamData) IsUsdtPair() bool {
	return strings.HasSuffix(s.Pairs, "USDT")
}

func (s streamData) Get24HourChange() float64 {
	now, _ := strconv.ParseFloat(s.Current, 64)
	open, _ := strconv.ParseFloat(s.Open, 64)
	return now/open - 1
}

type streamMessage struct {
	Stream string       `json:"stream"`
	Datas  []streamData `json:"data"`
}

var filterPaires = map[string]bool{
	"BTC":  true,
	"ETH":  true,
	"BUSD": true,
	"USDC": true,
	"TUSD": true,
}

func processMsg(msg []byte, sendNotify bool) (map[string]interface{}, bool) {
	var stream streamMessage
	err := json.Unmarshal(msg, &stream)
	if err != nil {
		fmt.Println("unmarshal msg failed ", err)
		return nil, false
	}
	if stream.Stream == "" {
		return nil, false
	}
	fmt.Println("pairs len:", len(stream.Datas))
	var parsed []parsedData
	for _, v := range stream.Datas {
		if !v.IsUsdtPair() || v.GetQuality() < 100000000 {
			continue
		}
		if filterPaires[v.Pairs[:len(v.Pairs)-4]] {
			continue
		}
		change := v.Get24HourChange()
		if change <= 0.03 {
			continue
		}
		parsed = append(parsed, parsedData{
			Pairs:   v.Pairs,
			Change:  change,
			Quality: v.GetQuality(),
			Current: v.GetCurrent(),
		})
	}
	// 先按照涨幅降序
	sort.Slice(parsed, func(i, j int) bool {
		return parsed[i].Change > parsed[j].Change
	})
	topIdx := len(parsed)
	if topIdx > len(parsed) {
		topIdx = len(parsed)
	}
	top := parsed[:topIdx]
	// 再按照成交量排序
	sort.Slice(top, func(i, j int) bool {
		return top[i].Quality > top[j].Quality
	})
	//for _, v := range top {
	//	fmt.Println(v)
	//}
	topIdx = 10
	if topIdx > len(top) {
		topIdx = len(top)
	}
	top = top[:topIdx]
	loc, _ := time.LoadLocation("Asia/Shanghai")
	timeStr := time.Now().In(loc).Format(`2006-01-02 15:04:05`)
	str := `<p style="font-size:1.1rem">` + timeStr + `</p>`
	markDownStr := "## " + timeStr + "\n\n```\n" + timeStr + "\n\n"
	for _, v := range top {
		val := fmt.Sprintf(`<p style="font-size:1.1rem">%s</p>`, v.String())
		str += val
		markDownStr += fmt.Sprintf("%s\n\n", v.String())
	}
	markDownStr += "\n\n```\n"
	fmt.Println("result:\n", str)
	fmt.Println("markDownStr:\n", markDownStr)
	if sendNotify {
		err = push2Server(mykey, markDownStr)
	}
	retMap := map[string]interface{}{
		"isBase64Encoded": false,
		"statusCode":      200,
		"headers":         map[string]string{"Content-Type": "text/html; charset=utf-8"},
		"body":            "<html><head><meta charset=\"UTF-8\"><body>" + str + "</body></html>",
	}
	return retMap, true
}

//go:embed mykey.txt
var mykey string

type parsedData struct {
	Pairs   string
	Change  float64
	Quality float64
	Current float64
}

func (p parsedData) String() string {
	currentFormat := "%.2f"
	if p.Current <= 0.01 {
		currentFormat = "%f"
	}
	return fmt.Sprintf("%s 价:"+currentFormat+" 幅:%.02f%% 额:%s", getUsdtParisName(p.Pairs), p.Current, p.Change*100, parseHumanReadableQuality(p.Quality))
}

func getUsdtParisName(pair string) string {
	return strings.Split(pair, "USDT")[0]
}

func parseHumanReadableQuality(val float64) string {
	if val < 10000 {
		return fmt.Sprintf("%.f", val)
	}
	if val < 1000000 {
		return fmt.Sprintf("%.02f 万", val/10000)
	}
	if val < 10000000 {
		return fmt.Sprintf("%.02f 百万", val/1000000)
	}
	if val < 100000000 {
		return fmt.Sprintf("%.02f 千万", val/10000000)
	}
	return fmt.Sprintf("%.02f 亿", val/100000000)
}
