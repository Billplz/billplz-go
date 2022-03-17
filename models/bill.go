package models

type Bill struct {
	CollectionId    string `json:"collection_id"`
	Email           string `json:"email"`
	Mobile          string `json:"mobile"`
	Name            string `json:"name"`
	Amount          int    `json:"amount"`
	CallbackUrl     string `json:"callback_url"`
	Description     string `json:"description"`
	DueAt           string `json:"due_at,omitempty"`
	RedirectUrl     string `json:"redirect_url,omitempty"`
	Deliver         bool   `json:"deliver,omitempty"`
	Reference1Label string `json:"reference_1_label,omitempty"`
	Reference1      string `json:"reference_1,omitempty"`
	Reference2Label string `json:"reference_2_label,omitempty"`
	Reference2      string `json:"reference_2,omitempty"`
}

type BillResponse struct {
	Bill
	ID         string      `json:"id"`
	Paid       bool        `json:"paid"`
	State      string      `json:"state"`
	PaidAmount int         `json:"paid_amount"`
	Url        string      `json:"url"`
	PaidAt     *CustomTime `json:"paid_at"`
}
