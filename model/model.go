package model

import (
	"time"
)

type Food struct {
	ID             int       `json:ID`
	Name           string    `json:Name`
	ExpirationDate time.Time `json:ExpirationDate`
	Position       int       `json:Position`
	PadNum         int       `json:PadNum`
	Count          int       `json:Count`
}

type Foods struct {
	Foods []Food `json:Foods`
}
