package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Receipt struct {
	Debug bool
	Port  string
}

func NewReceipt() (Receipt, error) {
	var r Receipt

	if len(os.Getenv("RECEIPT_DEBUG")) == 0 {
		err := godotenv.Load()
		if err != nil {
			return r, err
		}
	}

	r.Debug = os.Getenv("RECEIPT_DEBUG") == "TRUE"
	r.Port = os.Getenv("RECEIPT_PORT") // TODO: get default value

	return r, nil
}
