package repositories

import (
	"errors"
	"sync"

	"github.com/hilton-james/FetchExercise/internal/core/entities"
)

var (
	ErrDuplicateId = errors.New("receipt ID is already available")
	ErrNotFound    = errors.New("receipt not found")
)

type MemoryRepository struct {
	mux      sync.RWMutex
	receipts map[string]*entities.Receipt
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		receipts: make(map[string]*entities.Receipt),
	}
}

func (repo *MemoryRepository) Save(receipt *entities.Receipt) error {
	repo.mux.Lock()
	defer repo.mux.Unlock()

	if _, exists := repo.receipts[receipt.ID]; exists {
		return ErrDuplicateId
	}

	repo.receipts[receipt.ID] = receipt
	return nil
}

func (r *MemoryRepository) GetByID(id string) (*entities.Receipt, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	receipt, exists := r.receipts[id]
	if !exists {
		return nil, ErrNotFound
	}

	return receipt, nil
}
