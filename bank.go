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

func GetFpxBanks() (string) {
  banks := []models.FpxBank{
    {
      BankCode: "ABMB0212",
      BankName: "Alliance Bank",
    },
    {
      BankCode: "ABB0233",
      BankName: "Affin Bank",
    },
    {
      BankCode: "AMBB0209",
      BankName: "AmBank",
    },
    {
      BankCode: "BCBB0235",
      BankName: "CIMB Clicks",
    },
    {
      BankCode: "BIMB0340",
      BankName: "Bank Islam",
    },
    {
      BankCode: "BKRM0602",
      BankName: "Bank Rakyat",
    },
    {
      BankCode: "BMMB0341",
      BankName: "Bank Muamalat",
    },
    {
      BankCode: "BSN0601",
      BankName: "BSN",
    },
    {
      BankCode: "CIT0217",
      BankName: "Citibank Berhad",
    },
    {
      BankCode: "HLB0224",
      BankName: "Hong Leong Bank",
    },
    {
      BankCode: "HSBC0223",
      BankName: "HSBC Bank",
    },
    {
      BankCode: "KFH0346",
      BankName: "Kuwait Finance House",
    },
    {
      BankCode: "MB2U0227",
      BankName: "Maybank2u",
    },
    {
      BankCode: "MBB0227",
      BankName: "Maybank2E",
    },
    {
      BankCode: "MBB0228",
      BankName: "Maybank2E",
    },
    {
      BankCode: "OCBC0229",
      BankName: "OCBC Bank",
    },
    {
      BankCode: "PBB0233",
      BankName: "Public Bank",
    },
    {
      BankCode: "RHB0218",
      BankName: "RHB Now",
    },
    {
      BankCode: "SCB0216",
      BankName: "Standard Chartered",
    },
    {
      BankCode: "UOB0226",
      BankName: "UOB Bank",
    },
  }

	// s := string(body)
  return bank
}

// func CreatePayout(data models.Payout) (string) {
//   URL += "/api/v4/mass_payment_instructions"
//   requestBody, _ := json.Marshal(data)
//
//   client := &http.Client{}
//
//   req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(requestBody))
//   req.SetBasicAuth(APIKEY, "")
//   req.Header.Set("Content-type", "application/json")
//
//   resp, _ := client.Do(req)
//
//   body, err := ioutil.ReadAll(resp.Body)
//   if err != nil {
//     log.Fatalln(err)
//   }
//
//   s := string(body)
//   return s
// }
