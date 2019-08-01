package billplz

import (
  "fmt"
	"net/http"
  "io/ioutil"
  "bytes"
  "encoding/json"
  models "github.com/billplz/billplz-go/models"
)

func GetCollection(collectionId string) (models.CollectionResponse, models.Error) {
  client := &http.Client{}

  URL += fmt.Sprintf("/api/v4/collections/%s", collectionId)

  req, _ := http.NewRequest("GET", URL, nil)
  req.SetBasicAuth(APIKEY, "")

  resp, _ := client.Do(req)
  body, err := ioutil.ReadAll(resp.Body)

  obj, err1 := response(body, err, "collection")
  objString, _ := json.Marshal(obj)
  k := models.CollectionResponse{}
  json.Unmarshal(objString, &k)

  return k, err1
}

func CreateCollection(data models.Collection) (models.CollectionResponse, models.Error) {
  URL += "/api/v4/collections"
  requestBody, _ := json.Marshal(data)

  client := &http.Client{}

  req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(requestBody))
  req.SetBasicAuth(APIKEY, "")
  req.Header.Set("Content-type", "application/json")

  resp, _ := client.Do(req)
  body, err := ioutil.ReadAll(resp.Body)

  obj, err1 := response(body, err, "collection")
  objString, _ := json.Marshal(obj)
  k := models.CollectionResponse{}
  json.Unmarshal(objString, &k)

  return k, err1
}
