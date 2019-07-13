package billplz

import (
  "fmt"
	"net/http"
  "io/ioutil"
  // "log"
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

  stagingBanks := []models.FpxBank{
    {
      BankCode: "TEST0001*",
      BankName: "Test 0001",
    },
    {
      BankCode: "TEST0002*",
      BankName: "Test 0002",
    },
    {
      BankCode: "TEST0003*",
      BankName: "Test 0003",
    },
    {
      BankCode: "TEST0004*",
      BankName: "Test 0004",
    },
    {
      BankCode: "TEST0021*",
      BankName: "Test 0021",
    },
    {
      BankCode: "TEST0022*",
      BankName: "Test 0022",
    },
    {
      BankCode: "TEST0023*",
      BankName: "Test 0023",
    },
  }

  if ENVIRONMENT == "staging" {
    banks = append(banks, stagingBanks...)
  }

  b, err := json.Marshal(banks)
  if err != nil {
    panic(err)
  }

  return string(b)
}

func GetBankVerification(bankAccountNumber string) (string) {
  client := &http.Client{}

  URL += fmt.Sprintf("/api/v3/bank_verification_services/%s", bankAccountNumber)

  req, _ := http.NewRequest("GET", URL, nil)
  req.SetBasicAuth(APIKEY, "")

  resp, _ := client.Do(req)
  body, _ := ioutil.ReadAll(resp.Body)
	s := string(body)
  return s
}
