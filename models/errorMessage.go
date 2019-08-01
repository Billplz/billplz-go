package models

type Error struct {
  Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
  Type string `json:"type"`
  Message string `json:"message"`
}
