package billplz

import (
  "fmt"
	"net/http"
  "io/ioutil"
  "log"
  "bytes"
  "encoding/json"
  models "github.com/helmiruza/billplz-go/models"
)

func GetPayoutCollection(payoutCollectionId string) (string) {
  client := &http.Client{}

  URL += fmt.Sprintf("/api/v4/mass_payment_instruction_collections/%s", payoutCollectionId)

  req, _ := http.NewRequest("GET", URL, nil)
  req.SetBasicAuth(APIKEY, "")

  resp, _ := client.Do(req)
  body, _ := ioutil.ReadAll(resp.Body)
	s := string(body)
  return s
}

func CreatePayoutCollection(data models.PayoutCollection) (string) {
  URL += "/api/v4/mass_payment_instruction_collections"
  requestBody, _ := json.Marshal(data)

  client := &http.Client{}

  req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(requestBody))
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