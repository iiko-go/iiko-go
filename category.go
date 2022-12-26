package iiko

type Category struct {
	Id                    string `json:"id"`
	Name                  string `json:"name"`
	IsActive              bool   `json:"isActive"`
	IsDefaultForNewGuests bool   `json:"isDefaultForNewGuests"`
}
