package main

import (
	"binanceprice/run"
	"context"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

func main() {
	cloudfunction.Start(tencenCloudWrap)
	//result, err := run()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("result", result)
}

func tencenCloudWrap(ctx context.Context, event run.DefineEvent) (map[string]interface{}, error) {
	ret, err := run.Run(ctx, event)
	if err != nil {
		return nil, err
	}
	retMap := map[string]interface{}{
		"isBase64Encoded": false,
		"statusCode":      200,
		"headers":         map[string]string{"Content-Type": "text/html; charset=utf-8"},
		"body":            "<html><head><meta charset=\"UTF-8\"><body>" + ret.HtmlStr + "</body></html>",
	}
	return retMap, nil
}
