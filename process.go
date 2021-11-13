package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type streamData struct {
	Estr    string `json:"e"` // 事件类型
	E       int    `json:"E"` // 事件时间
	Pairs   string `json:"s"` // 交易对 BNBBTC  BTCUSDT
	Current string `json:"c"` // 最新成交价格
	Open    string `json:"o"` // 24小时前开始第一笔成交价格
	High    string `json:"h"` // 24小时内最高成交价
	Low     string `json:"l"` // 24小时内最低成交价
	Volumn  string `json:"v"` // 成交量
	Quality string `json:"q"` // 成交额
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

func processMsg(msg []byte) (string, bool) {
	var stream streamMessage
	err := json.Unmarshal(msg, &stream)
	if err != nil {
		fmt.Println("unmarshal msg failed ", err)
		return "", false
	}
	if stream.Stream == "" {
		return "", false
	}
	fmt.Println("pairs len:", len(stream.Datas))
	var parsed []parsedData
	for _, v := range stream.Datas {
		if !v.IsUsdtPair() {
			continue
		}
		change := v.Get24HourChange()
		if change < 0 {
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
	topIdx := 20
	if topIdx > len(parsed) {
		topIdx = len(parsed)
	}
	top := parsed[:topIdx]
	// 再按照成交量排序
	sort.Slice(top, func(i, j int) bool {
		return parsed[i].Quality > parsed[j].Quality
	})
	topIdx = 5
	if topIdx > len(top) {
		topIdx = len(top)
	}
	top = top[:topIdx]
	str := ""
	for idx, v := range top {
		str += fmt.Sprintf("%d. %s\n", idx+1, v.String())
	}
	fmt.Println("result:\n", str)
	return str, true
}

type parsedData struct {
	Pairs   string
	Change  float64
	Quality float64
	Current float64
}

func (p parsedData) String() string {
	return fmt.Sprintf("%s 涨幅:%.02f%% 成交额:%s", getUsdtParisName(p.Pairs), p.Change*100, parseHumanReadableQuality(p.Quality))
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
