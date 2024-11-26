package ports

import "github.com/hilton-james/FetchExercise/internal/core/entities"

type Repository interface {
	Save(receipt *entities.Receipt) error
	GetByID(id string) (*entities.Receipt, error)
}
