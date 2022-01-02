package bscscan

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetBscLatestDayTransaction get prev day transaction ,return timestamp,transactions,error
func GetBscLatestDayTransaction() (string, string, error) {
	url := "https://bscscan.com/chart/tx?output=csv"
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}
	arr := bytes.Split(body, []byte("\n"))
	if len(arr) < 3 {
		return "", "", err
	}
	val := arr[len(arr)-2]
	arr = bytes.Split(val, []byte(","))
	if len(arr) < 2 {
		return "", "", err
	}
	return strings.ReplaceAll(string(arr[1]), "\"", ""), strings.ReplaceAll(string(arr[2]), "\"", ""), nil
}
