package models

import "time"

type Product struct {
	Id       string   `json:"id"`
	Features []string `json:"features"`
}

type License struct {
	Id              string    `json:"id"`
	Product         string    `json:"product"`
	Features        []string  `json:"features"`
	Expiration      time.Time `json:"expiration"`
	ActivationLimit int       `json:"activation_limit"`
}

type Activation struct {
	Id      string `json:"id"`
	License string `json:"license"`
	Product string `json:"product"`
}
