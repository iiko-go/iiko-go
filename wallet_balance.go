package iiko

type WalletBalance struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Type    int    `json:"type"`
	Balance int    `json:"balance"`
}
