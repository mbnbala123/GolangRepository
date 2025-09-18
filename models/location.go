package models

type Location struct {
	AddressLine1 string `json:"AddressLine1"`
	City         string `json:"city"`
	State        string `json:"state"`
}
