package models

type FpxBank struct {
  BankCode string `json:"bank_code"`
  BankName string `json:"bank_name"`
}

type Bank struct {
  Name string `json:"name"`
  IdNo string `json:"id_no"`
  AccNo string `json:"acc_no"`
  Code string `json:"code"`
  Organization bool `json:"organization"`
}
