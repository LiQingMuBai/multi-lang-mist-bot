package trxfee

import "testing"

func TestTrxfeeClient_EnableTimesOrder(t *testing.T) {

	trxfee := NewTrxfeeClient("https://trxfee.io/", "733CA41D38EC1BA3476C5A860E357FC3", "E2BB3F45CCA1F6E15C828A5AEE918A4CE185D099A361E682C9A5A63611D0E154")

	trxfee.EnableTimesOrder("TY5t9HdU3h5ZT4LAw6E8ZN2jK297VL9999")
}
