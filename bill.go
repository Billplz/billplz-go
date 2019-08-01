package billplz

import (
  "fmt"
	"net/http"
  "io/ioutil"
  "bytes"
  "encoding/json"
  models "github.com/billplz/billplz-go/models"
)

func GetBill(billId string) (models.BillResponse, models.Error) {
  client := &http.Client{}

  URL += fmt.Sprintf("/api/v3/bills/%s", billId)

  req, _ := http.NewRequest("GET", URL, nil)
  req.SetBasicAuth(APIKEY, "")

  resp, _ := client.Do(req)
  body, err := ioutil.ReadAll(resp.Body)

  obj, err1 := response(body, err, "bill")
  objString, _ := json.Marshal(obj)
  k := models.BillResponse{}
  json.Unmarshal(objString, &k)

  return k, err1
}

func CreateBill(data models.Bill) (models.BillResponse, models.Error) {
  URL += "/api/v3/bills"
  requestBody, _ := json.Marshal(data)

  client := &http.Client{}

  req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(requestBody))
  req.SetBasicAuth(APIKEY, "")
  req.Header.Set("Content-type", "application/json")

  resp, _ := client.Do(req)
  body, err := ioutil.ReadAll(resp.Body)

  obj, err1 := response(body, err, "bill")
  objString, _ := json.Marshal(obj)
  k := models.BillResponse{}
  json.Unmarshal(objString, &k)

  return k, err1
}
