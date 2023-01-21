package usecases

import (
	"github.com/google/uuid"
	"github.com/mfitrahrmd/pastebin/internal/core/domain"
	"github.com/mfitrahrmd/pastebin/internal/core/ports"
	"github.com/mfitrahrmd/pastebin/internal/infrastructures/common"
	"github.com/mfitrahrmd/pastebin/internal/infrastructures/contentStorage"
	"time"
)

type pastebinUseCase struct {
	pastebinRepository     ports.PastebinRepository
	pastebinContentStorage contentStorage.PastebinContentStorage
}

// pastebin usecase constructor
func NewPastebinUseCase(pastebinRepository ports.PastebinRepository, pastebinContentStorage contentStorage.PastebinContentStorage) ports.PastebinUseCase {
	p := pastebinUseCase{
		pastebinRepository:     pastebinRepository,
		pastebinContentStorage: pastebinContentStorage,
	}

	return &p
}

// will save pastebin data into repo & the content into storage
// returning craeted pastebin data
// if shortlink is empty it will generate random shortlink
// pastebin can be expired with given length in minutes, 0 value for never expired
func (p *pastebinUseCase) Create(shortlink string, expiration uint, content []byte) (domain.Pastebin, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return domain.Pastebin{}, err
	}

	// generate random shortlink if its empty
	if shortlink == "" {
		shortlink = common.GenerateRandomString(6)
	}

	// save pastebin content into storage
	savedPastebinPath, err := p.pastebinContentStorage.Save(content)
	if err != nil {
		return domain.Pastebin{}, err
	}

	// pastebin data to be saved
	pb := domain.Pastebin{
		Base: domain.Base{
			Id:        newUUID.String(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Shortlink:                 shortlink,
		ExpirationLengthInMinutes: expiration,
		Path:                      savedPastebinPath,
	}

	// save pastebin data into repo
	err = p.pastebinRepository.Save(pb)
	if err != nil {
		return domain.Pastebin{}, err
	}

	return pb, nil
}

func (p *pastebinUseCase) Get(shortlink string) (domain.Pastebin, error) {
	//TODO implement me
	panic("implement me")
}

func (p *pastebinUseCase) List() ([]domain.Pastebin, error) {
	//TODO implement me
	panic("implement me")
}
