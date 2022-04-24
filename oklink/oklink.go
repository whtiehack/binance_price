package oklink

import (
	"crypto/tls"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//go:embed key.txt
var OKLINK_API_KEY string

func getApiKey() string {
	key := OKLINK_API_KEY[8:] + OKLINK_API_KEY[0:8]
	t := time.Now().UnixMilli()
	t = 1*t + 1111111111111
	tstr := strconv.Itoa(int(t)) + strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10))
	key = key + "|" + tstr
	return base64.StdEncoding.EncodeToString([]byte(key))
}

// curl 'https://www.oklink.com/api/explorer/v1/eth/info?t=1641093261881' \
//  -H 'Connection: keep-alive' \
//  -H 'sec-ch-ua: " Not;A Brand";v="99", "Microsoft Edge";v="97", "Chromium";v="97"' \
//  -H 'x-apiKey: LWIzMWUtNDU0Ny05Mjk5LWI2ZDA3Yjc2MzFhYmEyYzkwM2NjfDI3NTIyMDQzNzI5NzgyODQ=' \
//  -H 'x-cdn: https://static.oklink.com' \
//  -H 'devId: b02aa63f-6eca-4627-a7e0-7f80aab90ee7' \
//  -H 'Accept-Language: zh-CN' \
//  -H 'sec-ch-ua-mobile: ?0' \
//  -H 'Accept: application/json' \
//  -H 'x-utc: 8' \
//  -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.56 Safari/537.36 Edg/97.0.1072.41' \
//  -H 'App-Type: web' \
//  -H 'sec-ch-ua-platform: "Windows"' \
//  -H 'Sec-Fetch-Site: same-origin' \
//  -H 'Sec-Fetch-Mode: cors' \
//  -H 'Sec-Fetch-Dest: empty' \
//  -H 'Referer: https://www.oklink.com/zh-cn/eth'

func GetEthInfo() (Info, error) {
	key := getApiKey()
	url := "https://www.oklink.com/api/explorer/v1/eth/info"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Info{}, err
	}
	request.Header.Set("x-apiKey", key)
	request.Header.Set("x-cdn", "https://static.oklink.com")
	request.Header.Set("Accept-Language", "zh-CN")
	request.Header.Set("sec-ch-ua-mobile", "?0")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("x-utc", "8")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.56 Safari/537.36 Edg/97.0.1072.41")
	request.Header.Set("App-Type", "web")
	request.Header.Set("sec-ch-ua-platform", "Windows")
	request.Header.Set("Sec-Fetch-Site", "same-origin")
	request.Header.Set("Sec-Fetch-Mode", "cors")
	request.Header.Set("Sec-Fetch-Dest", "empty")
	request.Header.Set("Referer", "https://www.oklink.com/zh-cn/eth")
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	response, err := client.Do(request)
	if err != nil {
		return Info{}, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Info{}, err
	}
	var info Info
	err = json.Unmarshal(body, &info)
	return info, err
}
