package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RunScenarioResponse struct {
	RequestId string `json:"request_id"`
	Status    string `json:"status"`
}

func (s Service) RunScenario(token string, scenarioID string) error {
	url := fmt.Sprintf("https://api.iot.yandex.net/v1.0/scenarios/%s/actions", scenarioID)
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	resp := RunScenarioResponse{}
	if err = json.Unmarshal(body, &resp); err != nil {
		return err

	}
	if resp.Status != "ok" {
		return fmt.Errorf("run scenario failed: %s", scenarioID)
	}
	return nil
}
