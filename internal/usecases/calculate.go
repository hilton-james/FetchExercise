package usecases

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/hilton-james/FetchExercise/internal/core/entities"
)

var (
	ErrFailedToCovertToNumber = errors.New("failed to convert to number")
	ErrFailedToCovertToTime   = errors.New("failed to convert to Time")
)

func CalculatePoints(receipt *entities.Receipt) (int, error) {
	var points = 0

	points += len(strings.Map(func(r rune) rune {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			return r
		}
		return -1
	}, receipt.Retailer))

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, ErrFailedToCovertToNumber
	}

	if total == float64(int(total)) {
		points += 50
	}

	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			itemPrice, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, ErrFailedToCovertToNumber
			}

			points += int(math.Ceil(itemPrice * 0.2))
		}
	}

	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return 0, ErrFailedToCovertToTime
	}

	if date.Day()%2 != 0 {
		points += 6
	}

	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return 0, ErrFailedToCovertToTime
	}

	if (purchaseTime.Hour() == 14) || (purchaseTime.Hour() == 15 && purchaseTime.Minute() < 60) {
		points += 10
	}

	return points, nil
}
