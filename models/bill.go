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
	ID string	`json:"id"`
	CollectionId string	`json:"collection_id"`
	Paid bool	`json:"paid"`
	State string	`json:"state"`
	Amount float64	`json:"amount"`
	PaidAmount float64	`json:"paid_amount"`
	DueAt CustomTime`json:"due_at"`
	Email string	`json:"email"`
	Mobile string	`json:"mobile"`
	Name string	`json:"name"`
	Url string	`json:"url"`
	Reference1Label string	`json:"reference_1_label"`
	Reference1 *string	`json:"reference_1"`
	Reference2Label string	`json:"reference_2_label"`
	Reference2 *string	`json:"reference_2"`
	RedirectUrl *string	`json:"redirect_url"`
	CallbackUrl *string	`json:"callback_url"`
	Description string	`json:"description"`
	PaidAt *CustomTime	`json:"paid_at"`
}
