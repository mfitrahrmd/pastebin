package ports

import "github.com/mfitrahrmd/pastebin/internal/core/domain"

type PastebinUseCase interface {
	Create(shortlink string, expiration uint, content []byte) (domain.Pastebin, error)
	Get(shortlink string) (domain.Pastebin, error)
	GetAll() ([]domain.Pastebin, error)
}
