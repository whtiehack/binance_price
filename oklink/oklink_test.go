package oklink

import "testing"

func TestGetEthInfo(t *testing.T) {
	ethInfo, err := GetEthInfo()
	if err != nil {
		t.Error(err)
	}
	t.Log(ethInfo, ethInfo.Data.Transaction.TransactionValue24H)
}

func TestGetKey(t *testing.T) {
	k := getApiKey()
	t.Log(k)
}
