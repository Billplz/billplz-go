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

var ENVIRONTMENT = "";
var APIKEY = "";

func Init(e string, f string) {
	ENVIRONTMENT = e;
  APIKEY = f
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
  req.SetBasicAuth(APIKEY, "")

  response, _ := client.Do(req)
  bodyText, _ := ioutil.ReadAll(response.Body)
	s := string(bodyText)
  return s
}

// ei3a6mdl
