package main

import (
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
