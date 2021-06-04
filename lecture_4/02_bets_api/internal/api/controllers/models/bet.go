package models

// BetDto is a storage model representation of a bet with removed customerId.
type BetDto struct {
	Id                   string  `json:"id"`
	Status               string  `json:"status"`
	SelectionId          string  `json:"selection_id"`
	SelectionCoefficient float64 `json:"selection_id_coefficient"`
	Payment              float64 `json:"payment"`
	Payout               float64 `json:"payout"`
}
