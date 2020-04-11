package billplz

import (
  "encoding/json"

  models "github.com/billplz/billplz-go/models"
)

var ENVIRONMENT = ""
var APIKEY = ""
var URL = ""

const (
  ProductionUrl = "https://www.billplz.com"
	StagingUrl = "https://www.billplz-sandbox.com"
)

func Init(e string, f string) {
	ENVIRONMENT = e
  APIKEY = f

  if ENVIRONMENT == "production" {
    URL = ProductionUrl
  }

  if ENVIRONMENT == "staging" {
    URL = StagingUrl
  }
}

func objectFactory(objectType string) (interface{}) {
  switch (objectType) {
    case "collection":
      return models.CollectionResponse{}
    case "bill":
      return models.BillResponse{}
  }
  return nil
}

func response(body []byte, err error, objectType string) (interface{}, models.Error) {
  errorMessage := models.Error{}
  response := objectFactory(objectType)

  if err != nil {
    errorMessage.Error = models.ErrorDetail { Type: "unknown", Message: "unknown" }
    return response, errorMessage
  }

  json.Unmarshal(body, &errorMessage)

  if len(errorMessage.Error.Type) > 0 {
    return response, errorMessage
  }

  json.Unmarshal(body, &response)

  return response, errorMessage
}
