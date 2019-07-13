package models

type Payout struct {
  MassPaymentInstructionCollectionId string `json:"mass_payment_instruction_collection_id"`
  BankCode string `json:"bank_code"`
  BankAccountNumber string `json:"bank_account_number"`
  IdentityNumber string `json:"identity_number"`
  Name string `json:"name"`
  Description string `json:"description"`
  Total int `json:"total"`
}
