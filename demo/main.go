package main

import (
	"demo/service"
	apiNew "demo/service/api"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	api.HandleReq()
	api.HandleResp()
	apiNew.HandleError()
}