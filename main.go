package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"net/url"
	"os"
	"os/signal"
	"time"
)

func main() {
	cloudfunction.Start(run)
	//result, err := run()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("result", result)
}

func run() (string, error) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: "stream.yshyqxx.com", Path: "/stream"}
	fmt.Printf("connecting to %s\n", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic("dial:" + err.Error())
	}
	defer c.Close()

	done := make(chan struct{})
	var result string
	go func() {
		defer close(done)
		err := c.WriteMessage(websocket.TextMessage, []byte(`{"method":"SUBSCRIBE","params":["!miniTicker@arr@3000ms"],"id":1}`))
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
			fmt.Printf("recv: %s\n", message)
			msg, ok := processMsg(message)
			if ok {
				fmt.Println("process ended")
				result = msg
				break
			}
		}
	}()

	for {
		select {
		case <-done:
			return result, nil
		case <-interrupt:
			fmt.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fmt.Println("write close:", err)
				return result, err
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return result, err
		}
	}
}
