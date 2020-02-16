package main

import (
	"time"
)

const ip = "127.0.0.1"
const port = "9000"
const webSocketPort = "9001"
const socket = ip + ":" + port
const socketWebSocket = ip + ":" + webSocketPort
const readTimeoutRequest = 5 * time.Second
const writeTimeoutRequest = 10 * time.Second
const pathSeparator = "/"
const publicPath = "public" + pathSeparator
const templatePath = publicPath + "templates" + pathSeparator
const storagePath = "storage" + pathSeparator
const storageTmpFilePath = storagePath + "tmp" + pathSeparator
const kittenImgPath = publicPath + "img/kittens" + pathSeparator

func main() {
	go LoggerHandle()
	go runWebServer()
	go runWebSocketProcess()
	go KittenTaskHandle()
	select {}
}
