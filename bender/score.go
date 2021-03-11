package bender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/SECCDC/flexo/model"
	"github.com/lib/pq"
)

func Score(teams, targets []int, cat int, desc string, baseUrl string) (string, error) {
	url := fmt.Sprintf("%s/event", baseUrl)
	fmt.Println(url)

	var (
		targetsPqSlice pq.Int64Array
		teamsPqSlice   pq.Int64Array
	)

	for _, target := range targets {
		targetsPqSlice = append(targetsPqSlice, int64(target))
	}

	for _, team := range teams {
		teamsPqSlice = append(teamsPqSlice, int64(team))
	}

	event := model.Event{
		Targets:     targetsPqSlice,
		Teams:       teamsPqSlice,
		Category:    cat,
		Description: desc,
	}

	client := &http.Client{}

	body, err := json.Marshal(event)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
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
