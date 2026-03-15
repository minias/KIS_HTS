// internal/infrastructure/kis/approval.go
package kis

import (
	"bytes"
	"encoding/json"
	"net/http"

	"KIS_HTS/internal/config"
)

type ApprovalRequest struct {
	AppKey    string `json:"appkey"`
	AppSecret string `json:"secretkey"`
}

type ApprovalResponse struct {
	ApprovalKey string `json:"approval_key"`
}

// RequestApprovalKey websocket approval
func RequestApprovalKey(cfg *config.Config) (string, error) {

	body := ApprovalRequest{
		AppKey:    cfg.KIS.AppKey,
		AppSecret: cfg.KIS.AppSecret,
	}

	data, _ := json.Marshal(body)

	resp, err := http.Post(
		cfg.KIS.ApprovalURL,
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var result ApprovalResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.ApprovalKey, nil
}
