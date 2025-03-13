package models

// UserData represents a row from the Excel sheet
type UserData struct {
	Name   string `json:"name"`
	Date   string `json:"date"`
	Amount string `json:"amount"`
	Status string `json:"status"`
}
