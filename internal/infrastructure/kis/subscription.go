// internal/infrastructure/kis/subscription.go
package kis

import "encoding/json"

type SubscribeMessage struct {
	Header Header `json:"header"`
	Body   Body   `json:"body"`
}

type Header struct {
	ApprovalKey string `json:"approval_key"`
	CustType    string `json:"custtype"`
	TrType      string `json:"tr_type"`
	ContentType string `json:"content-type"`
}

type Body struct {
	Input Input `json:"input"`
}

type Input struct {
	TrID string `json:"tr_id"`
	Key  string `json:"tr_key"`
}

// BuildSubscribeMessage create subscription payload
func BuildSubscribeMessage(approvalKey, trID, symbol string) ([]byte, error) {

	msg := SubscribeMessage{
		Header: Header{
			ApprovalKey: approvalKey,
			CustType:    "P",
			TrType:      "1",
			ContentType: "utf-8",
		},
		Body: Body{
			Input: Input{
				TrID: trID,
				Key:  symbol,
			},
		},
	}

	return json.Marshal(msg)
}
