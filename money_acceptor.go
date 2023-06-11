package proto

// moneyin event data
type EventMoneyinData struct {
	// Amount of money inputted by the user
	Amount int64 `json:"amount"`

	// Optional currency field. Should contain "eur", "isk" etc.
	Currency string `json:"currency,omitempty"`
}
