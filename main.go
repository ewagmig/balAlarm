package main

import (
	"encoding/json"
	"fmt"
	"github.com/blinkbean/dingtalk"
	"io/ioutil"
	"net/http"
)

func main() {
	//c := cron.New()
	//c.AddFunc("CRON_TZ=Asia/Shanghai 0 23,5,11,17 * * *", SendNotification)
	//c.Start()
	fmt.Println("Begin to send notification to DingDing group!")
	SendNotification()
}

func SendNotification() {
	var dingToken = []string{"3b2cc7ddfb07fe4969f0d6eb1cb2c3e45d1289b05f5379d0bd0f4a41486f0c8d"}
	accountHT, err := GetBalanceByAPI()
	if err != nil {
		return
	}

	cli := dingtalk.InitDingTalk(dingToken, "+Balance")
	cli.SendTextMessage(accountHT)
}

type response struct {
	Status	string	`json:"status"`
	Message string	`json:"message"`
	Result  string	`json:"result"`
}
func GetBalanceByAPI() (string, error) {
	resp, err := http.Get("https://api.hecoinfo.com/api?module=account&action=balance&address=0x4188d0da3a993f77bbbb57e15c16dccf035c1ef8&tag=latest&apikey=YourApiKeyToken")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return "", err
	}

	var rep response
	err = json.Unmarshal(respBody, &rep)
	if err != nil {
		return "", err
	}

	baStr := rep.Result
	ba18 := baStr[:len(baStr)-18]
	BaHT := ba18 + "HT"
	return BaHT, nil
}