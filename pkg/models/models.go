package models

import "time"

type Product struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Entitlements []string `json:"features"`
}

type License struct {
	Id              string    `json:"id"`
	Product         string    `json:"product"`
	Expiration      time.Time `json:"expiration"`
	ActivationLimit int       `json:"activation_limit"`
	Entitlements    []string  `json:"features"`
}

type Activation struct {
	Id      string `json:"id"`
	License string `json:"license"`
	Product string `json:"product"`
}
