package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func push2Server(key, msg string) error {
	resp, err := http.PostForm("https://sctapi.ftqq.com/"+key+".send", url.Values{
		"title": []string{"coin data"},
		"desp":  []string{msg},
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	val, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("server jiang resp:", string(val))
	return nil
}

func push2WxPusher(key, group, msg string) error {
	summary := msg
	if len(summary) > 100 {
		summary = summary[:100]
	}
	b, _ := json.Marshal(map[string]interface{}{
		"appToken":    key,
		"content":     msg,
		"contentType": 2,
		"topicIds":    []string{group},
		"summary":     summary,
	})
	resp, err := http.Post("http://wxpusher.zjiecode.com/api/send/message", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	val, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("push2WxPusher resp:", string(val))
	return nil
}
