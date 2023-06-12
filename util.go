package proto

import "encoding/json"

func GetScanData(r json.RawMessage) (*EventScanData, error) {
	var e EventScanData
	err := json.Unmarshal(r, &e)
	return &e, err
}

func GetMoneyinData(r json.RawMessage) (*EventMoneyinData, error) {
	var e EventMoneyinData
	err := json.Unmarshal(r, &e)
	return &e, err
}
