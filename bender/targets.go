package bender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/SECCDC/flexo/model"
)

func GetTargets(baseUrl string) ([]model.Target, error) {
	var (
		targets   []model.Target
		bodyBytes []byte
	)
	url := fmt.Sprintf("%s/targets", baseUrl)
	fmt.Println(url)

	target := model.Target{}
	client := &http.Client{}

	body, err := json.Marshal(target)
	if err != nil {
		return targets, err
	}

	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(body))
	if err != nil {
		return targets, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return targets, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return targets, fmt.Errorf("Bad status code: %d", resp.StatusCode)
	}

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &targets)

	return targets, err
}
