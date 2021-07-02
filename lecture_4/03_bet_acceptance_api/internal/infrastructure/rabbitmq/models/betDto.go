package models

// BetDto represents a DTO bet.
type BetDto struct {
	Id                   string  `json:"id"`
	CustomerId           string  `json:"customerId"`
	SelectionId          string  `json:"selectionId"`
	SelectionCoefficient float64 `json:"selectionCoefficient"`
	Payment              float64 `json:"payment"`
}
