package bender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/SECCDC/flexo/model"
)

func GetTeams(baseUrl string) (string, error) {
	url := fmt.Sprintf("%s/teams", baseUrl)
	fmt.Println(url)

	team := model.Team{}
	client := &http.Client{}

	body, err := json.Marshal(team)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return string(bodyBytes), nil
	}
	return "", fmt.Errorf("Returned bad HTTP status: %d", resp.StatusCode)
}
