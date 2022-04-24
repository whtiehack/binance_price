package run

import (
	"binanceprice/bscscan"
	"binanceprice/oklink"
	"context"
	"crypto/tls"
	_ "embed"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

type DefineEvent struct {
	Type string
	// test event define
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

type RunResult struct {
	HtmlStr     string
	MarkdownStr string
}

//go:embed mykey.txt
var mykey string

func Run(_ context.Context, event DefineEvent) (RunResult, error) {
	fmt.Printf("event: %#v\n", event)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: "stream.yshyqxx.com", Path: "/stream"}
	fmt.Printf("connecting to %s\n", u.String())

	dialer := &websocket.Dialer{TLSClientConfig: &tls.Config{RootCAs: nil, InsecureSkipVerify: true}}
	c, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		panic("dial:" + err.Error())
	}
	defer c.Close()
	var htmlStr, markDownStr string
	var ethValue24h int
	var wg sync.WaitGroup
	wg.Add(3)
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
		for i := 0; i < 3; i++ {
			ethInfo, err := oklink.GetEthInfo()
			if err != nil {
				fmt.Println("get eth info error:", err)
				return
			}
			ethValue24h = int(ethInfo.Data.Transaction.TransactionValue24H)
			if ethValue24h != 0 {
				break
			}
			fmt.Println("get eth value 24h error, retry...", i)
		}
	}()
	var bscTimeStamp, bscTransactions string
	go func() {
		defer wg.Done()
		var err error
		count := 0
		for count < 4 {
			count++
			bscTimeStamp, bscTransactions, err = bscscan.GetBscLatestDayTransaction()
			if err != nil || bscTimeStamp == "" {
				fmt.Println("get bsc latest day transaction error:", err)
				continue
			}
			break
		}
	}()
	wg.Wait()
	htmlStr += `<p style="font-size:1.1rem">24h 链上交易量</p>`
	ethVal := comma(strconv.Itoa(ethValue24h))
	htmlStr += `<p style="font-size:1.1rem">` + ethVal + `ETH</p>`
	markDownStr += "\n\n```\n24h 链上交易量\n" + ethVal + "ETH\n```\n"

	if bscTimeStamp != "" {
		t, _ := strconv.Atoi(bscTimeStamp)
		fmt.Println("bsc timestamp:", bscTimeStamp, bscTransactions, time.Unix(int64(t), 0).Format("2006-01-02"))
		htmlStr += `<p style="font-size:1.1rem">` + time.Unix(int64(t), 0).Format("2006-01-02") + ` BSC链上交易数</p>`
		bscTransactions = comma(bscTransactions)
		htmlStr += `<p style="font-size:1.1rem">` + bscTransactions + `</p>`
		markDownStr += "\n\n```\n" + time.Unix(int64(t), 0).Format("2006-01-02") + " BSC链上交易数\n" + bscTransactions + "\n```\n"
	}

	sendNotify := event.Type == "Timer"
	if sendNotify && htmlStr != "" {
		e := push2WxPusher(mykey, "5580", htmlStr)
		if e != nil {
			fmt.Println("push to server error:", e)
		}
	}
	fmt.Println("result:\n", htmlStr)
	fmt.Println("markDownStr:\n", markDownStr)
	return RunResult{
		HtmlStr:     htmlStr,
		MarkdownStr: markDownStr,
	}, nil
}
