module binanceprice

go 1.18

require (
	github.com/aliyun/fc-runtime-go-sdk v0.0.4
	github.com/gorilla/websocket v1.4.2
	github.com/tencentyun/scf-go-lib v0.0.0-20200624065115-ba679e2ec9c9
)

require (
    huaweicloud.com/go-runtime v0.0.0-00010101000000-000000000000
)

replace (
    huaweicloud.com/go-runtime => ./go-runtime
)