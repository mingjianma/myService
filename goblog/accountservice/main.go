package main

import (
    "fmt"
    "github.com/callistaenterprise/goblog/accountservice/service" // 新增代码
    "github.com/callistaenterprise/goblog/accountservice/dbclient"
)

var appName = "accountservice"

func main() {
    fmt.Printf("Starting %v\n", appName)
    initializeBoltClient()
    service.StartWebServer("6767") // 新增代码
}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
    service.DBClient = &dbclient.BoltClient{}
    service.DBClient.OpenBoltDb()
    service.DBClient.Seed()
}
