package main

import (
	"strconv"
	"testing"
)

func TestGetBalanceByAPI(t *testing.T) {
	ba, err := GetBalanceByAPI()
	if err != nil {
		t.Error(err)
	}
	t.Log(ba)
}

func TestSplitStr(t *testing.T) {
	str := "3549456498065706531967"
	s18 := str[:len(str)-18]
	t.Log(s18)
}

func TestSendNotification(t *testing.T) {
	SendNotification()
}

func TestStrConv(t *testing.T) {
	str := "3000HT"
	amStr := str[:len(str)-2]
	am, err := strconv.Atoi(amStr)
	if err != nil {
		t.Error(err)
	}
	t.Log(am)
}