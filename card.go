package iiko

type Card struct {
	Id          string `json:"id"`
	Track       string `json:"track"`
	Number      string `json:"number"`
	ValidToDate string `json:"validToDate"`
}
