package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Receipt struct {
	Debug bool
}

func NewReceipt() (Receipt, error) {
	var r Receipt

	if len(os.Getenv("DEBUG")) == 0 {
		err := godotenv.Load()
		if err != nil {
			return r, err
		}
	}

	r.Debug = os.Getenv("DEBUG") == "TRUE"

	return r, nil
}
