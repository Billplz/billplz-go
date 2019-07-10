package billplz

import (
  "fmt"
	"net/http"
  "io/ioutil"
)

const (
  production_url = "https://billplz.com"
	staging_url = "https://billplz-staging.herokuapp.com"
)

func basicAuth(username, password string) string {
  auth := username + ":" + password
  return auth
}

var ENVIRONTMENT = "";

func Init(e string) {
	ENVIRONTMENT = e;
}

func GetCollection(collectionId string) (string) {
  env := ENVIRONTMENT // e.A
  url := ""
  client := &http.Client{}

  if env == "production" {
    url = production_url
  }

  if env == "staging" {
    url = staging_url
  }

  url += fmt.Sprintf("/api/v3/collections/%s", collectionId)

  req, _ := http.NewRequest("GET", url, nil)
  req.SetBasicAuth("69da23bf-da10-4fda-814d-3ad970035d38", "")

  response, _ := client.Do(req)
  bodyText, _ := ioutil.ReadAll(response.Body)
	s := string(bodyText)
  return s
}

// ei3a6mdl
