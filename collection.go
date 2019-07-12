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

func Init(e string, f string) {
	constants.ENVIRONTMENT = e
  constants.APIKEY = f

  if constants.ENVIRONTMENT == "production" {
    constants.URL = constants.ProductionUrl
  }

  if constants.ENVIRONTMENT == "staging" {
    constants.URL = constants.StagingUrl
  }
}

func GetCollection(collectionId string) (string) {
  client := &http.Client{}

  constants.URL += fmt.Sprintf("/api/v4/collections/%s", collectionId)

  req, _ := http.NewRequest("GET", constants.URL, nil)
  req.SetBasicAuth(constants.APIKEY, "")

  resp, _ := client.Do(req)
  body, _ := ioutil.ReadAll(resp.Body)
	s := string(body)
  return s
}

func CreateCollection(data models.Collection) (string) {
  constants.URL += "/api/v4/collections"
  requestBody, _ := json.Marshal(data)

  client := &http.Client{}

  req, _ := http.NewRequest("POST", constants.URL, bytes.NewBuffer(requestBody))
  req.SetBasicAuth(constants.APIKEY, "")
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
