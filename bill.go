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

  return response(body, err)
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

  return response(body, err)
}

func response(body []byte, err error) (models.BillResponse, models.Error) {
  errorMessage := models.Error{}

  if err != nil {
    errorMessage.Error = models.ErrorDetail { Type: "unknown", Message: "unknown" }
    return models.BillResponse{},errorMessage
  }

  response := models.BillResponse{}
  json.Unmarshal(body, &errorMessage)

  if len(errorMessage.Error.Type) > 0 {
    return response, errorMessage
  }

  json.Unmarshal(body, &response)

  return response, errorMessage
}
