package entities

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

var (
	retailerRegex        = regexp.MustCompile(`^[\w\s\-&]+$`)
	purchaseDateRegex    = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	purchaseTimeRegex    = regexp.MustCompile(`^\d{2}:\d{2}$`)
	totalRegex           = regexp.MustCompile(`^\d+\.\d{2}$`)
	itemDescriptionRegex = regexp.MustCompile(`^[\w\s\-\+]+$`)
)

type Receipt struct {
	ID           string `json:"id"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

func (r *Receipt) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Retailer, validation.Required, validation.Match(retailerRegex)),
		validation.Field(&r.PurchaseDate, validation.Required, validation.Match(purchaseDateRegex)),
		validation.Field(&r.PurchaseTime, validation.Required, validation.Match(purchaseTimeRegex)),
		validation.Field(&r.Total, validation.Required, validation.Match(totalRegex)),
		validation.Field(&r.Items, validation.Required, validation.Length(1, 0)),
	)
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

func (i *Item) Validate() error {
	return validation.ValidateStruct(i,
		validation.Field(&i.ShortDescription, validation.Required, validation.Match(itemDescriptionRegex)),
		validation.Field(&i.Price, validation.Required, validation.Match(totalRegex)),
	)
}
