package repositories

import (
	"errors"
	"github.com/mfitrahrmd/pastebin/internal/core/domain"
)

type inMemoryPastebinRepository struct {
	pastebin map[string]domain.Pastebin
}

func NewInMemoryPastebinRepository() *inMemoryPastebinRepository {
	pr := inMemoryPastebinRepository{
		pastebin: map[string]domain.Pastebin{},
	}

	return &pr
}

func (i *inMemoryPastebinRepository) Save(pastebin domain.Pastebin) error {
	i.pastebin[pastebin.Shortlink] = pastebin

	return nil
}

func (i *inMemoryPastebinRepository) Get(shortlink string) (domain.Pastebin, error) {
	p, ok := i.pastebin[shortlink]
	if !ok {
		return domain.Pastebin{}, errors.New("pastebin not found")
	}

	return p, nil
}

func (i *inMemoryPastebinRepository) List() ([]domain.Pastebin, error) {
	pastebin := []domain.Pastebin{}

	for _, p := range i.pastebin {
		pastebin = append(pastebin, p)
	}

	return pastebin, nil
}
