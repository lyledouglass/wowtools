package utilities

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	WowGameApi = "https://us.api.blizzard.com/data/wow/"
)

func GetTokenPrice(token string) int {

	url := WowGameApi + "token/index?namespace=dynamic-us&locale=en_US&access_token=" + token
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatalln(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	type tokenApiData struct {
		Price int `json:"price"`
	}
	var response tokenApiData
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		return 0
	}

	return response.Price
	// fmt.Println(response.Price)
}
