package main

import (
	"encoding/json"
	"net"
	"os"
	"strings"
	"log"
)

type resJSON struct {
	ID     int         `json:"id"`
	Error  interface{} `json:"error"`
	Result []string    `json:"result"`
}

var valuesETH = []string{
	"miner version",
	"running time, in minutes",
	"total ETH hashrate in MH/s",
	"number of ETH shares",
	"number of ETH rejected shares",
	"detailed ETH hashrate for all GPUs",
	"total DCR hashrate in MH/s",
	"number of DCR shares",
	"number of DCR rejected shares",
	"detailed DCR hashrate for all GPUs",
	"Temperature for all GPUs",
	"Fan speed(%) pairs ",
	"current mining pool. For dual mode, there will be two pools here",
	"number of ETH invalid shares",
	"number of ETH pool switches",
	"number of DCR invalid shares",
	"number of DCR pool switches",
}
var valuesZEC = []string{
	"miner version",
	"running time, in minutes",
	"total ETH hashrate in MH/s",
	"number of ETH shares",
	"number of ETH rejected shares",
	"detailed ETH hashrate for all GPUs",
	"total DCR hashrate in MH/s",
	"number of DCR shares",
	"number of DCR rejected shares",
	"detailed DCR hashrate for all GPUs",
	"Temperature for all GPUs",
	"Fan speed(%) pairs ",
	"current mining pool. For dual mode, there will be two pools here",
	"number of ETH invalid shares",
	"number of ETH pool switches",
	"number of DCR invalid shares",
	"number of DCR pool switches",
}
var serverAddr = "172.16.1.170:4028"

func getInfoFromMiner() []string {
	var minerInfo = make([]string, len(valuesZEC))
	strEcho := `{"id":0,"jsonrpc":"2.0","method":"miner_getstat1"}`
//	serverAddr := "172.16.1.170:4028"
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		log.Fatal("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	reply := make([]byte, 1024)

	readLen, err := conn.Read(reply)
	if err != nil {
		log.Fatal("Write to server failed:", err.Error())
		os.Exit(1)
	}

	var replyJSON resJSON

	if err := json.Unmarshal(reply[:readLen], &replyJSON); err != nil {
		panic(err)
	}
	if len(replyJSON.Result) <= len(valuesETH){
		if strings.Contains(replyJSON.Result[0], "ZEC") {

			iterator := 0
			for _, result := range replyJSON.Result {
				if strings.Contains(result, ";") {
					for _, subresult := range strings.Split(result, ";") {
						minerInfo[iterator] = valuesZEC[iterator] + ": <b>" + subresult + "</b>"
						iterator++
					}
				} else {
					minerInfo[iterator] = valuesZEC[iterator] + ": <b>" + result + "</b>"
					println(iterator)
					iterator++
				}

			}
			return minerInfo
		}

		if strings.Contains(replyJSON.Result[0], "ETH") {
			iterator := 0
			for _, result := range replyJSON.Result {
				if strings.Contains(result, ";") {
					for _, subresult := range strings.Split(result, ";") {
						minerInfo[iterator] = valuesETH[iterator] + ": <b>" + subresult + "</b>"
						println(iterator)
						iterator++
					}
				} else {
					minerInfo[iterator] = valuesETH[iterator] + ": <b>" + result + "</b>"
					println(iterator)
					iterator++
				}

			}
			return minerInfo
		}
	}



	return minerInfo
}
