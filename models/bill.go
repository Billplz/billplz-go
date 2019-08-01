package models

type Bill struct {
  CollectionId string `json:"collection_id"`
  Email string `json:"email"`
  Mobile string `json:"mobile"`
  Name string `json:"name"`
  Amount int `json:"amount"`
  CallbackUrl string `json:"callback_url"`
  Description string `json:"description"`
  DueAt string `json:"due_at"`
  RedirectUrl string `json:"redirect_url"`
  Deliver bool `json:"deliver"`
  Reference1Label string `json:"reference_1_label"`
  Reference1 string `json:"reference_1"`
  Reference2Label string `json:"reference_2_label"`
  Reference2 string `json:"reference_2"`
}

type BillResponse struct {
  
}
