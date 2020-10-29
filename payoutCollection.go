package billplz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	models "github.com/billplz/billplz-go/models"
)

func GetPayoutCollection(payoutCollectionId string) string {
	client := &http.Client{}

	url := fmt.Sprintf(URL+"/api/v4/mass_payment_instruction_collections/%s", payoutCollectionId)

	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(APIKEY, "")

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	s := string(body)
	return s
}

func CreatePayoutCollection(data models.PayoutCollection) string {
	url := URL + "/api/v4/mass_payment_instruction_collections"
	requestBody, _ := json.Marshal(data)

	client := &http.Client{}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.SetBasicAuth(APIKEY, "")
	req.Header.Set("Content-type", "application/json")

	resp, _ := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	s := string(body)
	return s
}

// ei3a6mdl
