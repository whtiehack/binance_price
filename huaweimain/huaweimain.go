package main

import (
	"binanceprice/run"
	context2 "context"
	"encoding/json"
	"fmt"
	"huaweicloud.com/go-runtime/events/apig"
	"huaweicloud.com/go-runtime/events/timer"
	"huaweicloud.com/go-runtime/go-api/context"
	"huaweicloud.com/go-runtime/pkg/runtime"
)

// cron
// 0 0 9,15,21 * * * *

func Apig(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var apigEvent apig.APIGTriggerEvent
	var timerEvent timer.TimerTriggerEvent
	err := json.Unmarshal(payload, &apigEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	isTimer := false
	if apigEvent.HttpMethod == "" {
		fmt.Println("HttpMethod is empty")
		err = json.Unmarshal(payload, &timerEvent)
		if err != nil {
			fmt.Println("Unmarshal failed")
			return "invalid data", err
		}
		fmt.Println("timerEvent: ", timerEvent.String())
		isTimer = true
	}
	de := run.DefineEvent{}
	if isTimer {
		de.Type = "Timer"
	}
	ret, err := run.Run(context2.Background(), de)
	if err != nil {
		fmt.Println("Run failed")
		if isTimer {
			return "invalid data", err
		}
		apigResp := apig.APIGTriggerResponse{
			Body: apigEvent.String() + "  error:" + err.Error(),
			Headers: map[string]string{
				"content-type": "text/html;charset=utf-8",
			},
			StatusCode: 200,
		}
		return apigResp, nil
	}
	if isTimer {
		return ret.MarkdownStr, nil
	}
	ctx.GetLogger().Logf("payload:%s", apigEvent.String())
	apigResp := apig.APIGTriggerResponse{
		Body: ret.HtmlStr,
		Headers: map[string]string{
			"content-type": "text/html;charset=utf-8",
		},
		StatusCode: 200,
	}
	return apigResp, nil
}

func main() {
	runtime.Register(Apig)
}
