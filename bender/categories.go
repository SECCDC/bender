package bender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/SECCDC/flexo/model"
)

func GetCategories(baseUrl string) ([]model.Category, error) {
	var (
		teams     []model.Category
		bodyBytes []byte
	)
	url := fmt.Sprintf("%s/categories", baseUrl)
	fmt.Println(url)

	team := model.Category{}
	client := &http.Client{}

	body, err := json.Marshal(team)
	if err != nil {
		return teams, err
	}

	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(body))
	if err != nil {
		return teams, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return teams, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return teams, fmt.Errorf("Bad status code: %d", resp.StatusCode)
	}

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &teams)

	return teams, err
}
