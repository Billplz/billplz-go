package billplz

var ENVIRONTMENT = ""
var APIKEY = ""
var URL = ""

const (
  ProductionUrl = "https://billplz.com"
	StagingUrl = "https://billplz-staging.herokuapp.com"
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
