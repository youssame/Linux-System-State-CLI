package network

import "testing"

func TestDisplayTitle(t *testing.T) {
	res := getNetworkMockedInfos().DisplayTitle()
	exp := `
Network Information:
--------------------`
	if res != exp {
		t.Errorf("DisplayTitle failed, expected %s, got %s", exp, res)
	}
}
func TestDisplayInfo(t *testing.T) {
	res := getNetworkMockedInfos().DisplayInfo()
	exp := `IP Address      : IP
Data Sent       : 1 MB
Data Received   : 2 GB`
	if res != exp {
		t.Errorf("got %s, want %s", res, exp)
	}
}

func getNetworkMockedInfos() NetworkInfo {
	return NetworkInfo{
		ip:            "IP",
		sent_data:     1,
		received_data: 2,
	}
}
