package main

import (
	"binanceprice/oklink"
	"context"
	_ "embed"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"sync"
)

func main() {
	cloudfunction.Start(run)
	//result, err := run()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("result", result)
}

type DefineEvent struct {
	Type string
	// test event define
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

//go:embed mykey.txt
var mykey string

func run(_ context.Context, event DefineEvent) (map[string]interface{}, error) {
	fmt.Printf("event: %#v\n", event)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: "stream.yshyqxx.com", Path: "/stream"}
	fmt.Printf("connecting to %s\n", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic("dial:" + err.Error())
	}
	defer c.Close()
	var htmlStr, markDownStr string
	var ethValue24h int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := c.WriteMessage(websocket.TextMessage, []byte(`{"method":"SUBSCRIBE","params":["!ticker@arr@3000ms"],"id":1}`))
		if err != nil {
			fmt.Println("write:", err)
			return
		}
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				fmt.Println("read:", err)
				return
			}
			var ok bool
			htmlStr, markDownStr, ok = processMsg(message)
			if ok {
				fmt.Println("binance prices process ended")
				break
			}
			fmt.Printf("recv: %s\n", message)
		}
	}()
	go func() {
		defer wg.Done()
		ethInfo, err := oklink.GetEthInfo()
		if err != nil {
			fmt.Println("get eth info error:", err)
			return
		}
		ethValue24h = int(ethInfo.Data.Transaction.TransactionValue24H)
	}()
	wg.Wait()
	htmlStr += `<p style="font-size:1.1rem">24h 链上交易量</p>`
	ethVal := comma(strconv.Itoa(ethValue24h))
	htmlStr += `<p style="font-size:1.1rem">` + ethVal + `ETH</p>`
	markDownStr += "\n\n```\n24h 链上交易量\n" + ethVal + "ETH\n```\n"

	sendNotify := event.Type == "Timer"
	if sendNotify && markDownStr != "" {
		e := push2Server(mykey, markDownStr)
		if e != nil {
			fmt.Println("push to server error:", e)
		}
	}
	fmt.Println("result:\n", htmlStr)
	fmt.Println("markDownStr:\n", markDownStr)
	retMap := map[string]interface{}{
		"isBase64Encoded": false,
		"statusCode":      200,
		"headers":         map[string]string{"Content-Type": "text/html; charset=utf-8"},
		"body":            "<html><head><meta charset=\"UTF-8\"><body>" + htmlStr + "</body></html>",
	}
	return retMap, nil
}
