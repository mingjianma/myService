package main

import (
    "fmt"
    "myService/crawlads/adsservice/service" // 新增代码
    //"myService/crawlads/adsservice/dbclient"
)

var appName = "accountservice"

func main() {
    fmt.Printf("Starting %v\n", appName)
    service.StartWebServer("6767") // 新增代码
}