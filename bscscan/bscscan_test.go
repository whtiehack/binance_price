package bscscan

import "testing"

func TestGetBscLatestDayTransaction(t *testing.T) {
	ts, val, err := GetBscLatestDayTransaction()
	if err != nil {
		t.Error(err)
	}
	if ts == "" {
		t.Error("timestamp is empty")
	}
	t.Log(ts, val)
}

