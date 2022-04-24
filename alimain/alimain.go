package main

import (
	"binanceprice/run"
	"context"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	req.Body.Close()
	ret, err := run.Run(ctx, run.DefineEvent{})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("runError: " + err.Error()))
		return nil
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/html;charset=utf-8")

	w.Write([]byte(ret.HtmlStr))
	return nil
}

func main() {
	fc.StartHttp(HandleHttpRequest)
}
