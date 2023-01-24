package ports

import "github.com/mfitrahrmd/pastebin/internal/core/domain"

type PastebinRepository interface {
	Save(pastebin domain.Pastebin) error
	Get(shortlink string) (domain.Pastebin, error)
	List() ([]domain.Pastebin, error)
}

type PastebinContentRepository interface {
	Save(content []byte) (string, error)
}
