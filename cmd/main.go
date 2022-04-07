package main

import (
	"fmt"
	"unsafe"
)

type NlpApiResult struct {
	AccountId     uint32      `json:"account_id"`
	AppId         uint32      `json:"app_id"`
	FunctionId    uint32      `json:"function_id"`
	EngineId      uint32      `json:"engine_id"`
	InputTextList []string    `json:"input_text_list"`
	ApiResult     interface{} `json:"api_result"`
	MsgId         string      `json:"msg_id"`
}

type aa struct {
	name *[]string   // 24
	s    string      // 16
	age  interface{} // 16
}

func main() {
	var a aa
	fmt.Println(unsafe.Sizeof(a))
}
