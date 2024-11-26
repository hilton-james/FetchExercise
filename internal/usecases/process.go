package usecases

import (
	"github.com/google/uuid"
	"github.com/hilton-james/FetchExercise/internal/core/entities"
	"github.com/hilton-james/FetchExercise/internal/ports"
)

func ProcessReceipt(repo ports.Repository, receipt *entities.Receipt) (string, error) {
	// TODO: consider expiration time if need be
	receipt.ID = uuid.New().String()

	// TODO: here we can check if it's duplicated
	if err := repo.Save(receipt); err != nil {
		return "", err
	}

	return receipt.ID, nil
}
