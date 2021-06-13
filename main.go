package main

import (
	"encoding/json"
	"fmt"
	"github.com/blinkbean/dingtalk"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	//c := cron.New()
	//c.AddFunc("CRON_TZ=Asia/Shanghai 0 23,5,11,17 * * *", SendNotification)
	//c.Start()
	fmt.Println("Begin to send notification to DingDing group!")
	SendNotification()
}

func SendNotification() {
	var dingToken = []string{"e742a486c292a0e1b82bf6a3e975e832f38ee47210c1dee81727a177e7dcc177"}
	accountHT, err := GetBalanceByAPI()
	if err != nil {
		return
	}
	amountNum := accountHT[:len(accountHT)-2]
	amount, err := strconv.Atoi(amountNum)
	if err != nil {
		return
	}
	if amount <= 2000{
		content := "Balance is lower than 2000HT, Please refund immediately!!!"
		//@周李
		clia := dingtalk.InitDingTalk(dingToken, "-Alarm")
		mobiles := []string{"13488858435"}
		clia.SendTextMessage(content, dingtalk.WithAtMobiles(mobiles))
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