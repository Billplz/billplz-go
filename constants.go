package billplz

var ENVIRONTMENT = ""
var APIKEY = ""
var URL = ""

const (
  ProductionUrl = "https://billplz.com"
	StagingUrl = "https://billplz-staging.herokuapp.com"
)

func Init(e string, f string) {
	ENVIRONTMENT = e
  APIKEY = f

  if ENVIRONTMENT == "production" {
    URL = ProductionUrl
  }

  if ENVIRONTMENT == "staging" {
    URL = StagingUrl
  }
}
