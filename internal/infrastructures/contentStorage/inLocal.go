package contentStorage

import (
	"fmt"
	"github.com/mfitrahrmd/pastebin/internal/infrastructures/common"
	"os"
	"strconv"
	"time"
)

const (
	defaultFileLocation = "."
	defaultFolderName   = "content"
	defaultFileExt      = ".txt"
)

var defaultConfig = Config{
	PastebinContentFileLocation: defaultFileLocation,
	PastebinContentFolderName:   defaultFolderName,
	FileExt:                     defaultFileExt,
	FileNamePrefix:              "",
	FileNameSuffix:              "",
}

type pastebinContentLocalStorage struct {
	config Config
}

func NewPastebinContentLocalStorage(config ...Config) PastebinContentStorage {
	s := pastebinContentLocalStorage{}

	err := s.initialize(config...)
	if err != nil {
		panic(fmt.Sprintf("error creating pastebin content local storage instance : %v", err))
	}

	return &s
}

func (p *pastebinContentLocalStorage) initialize(config ...Config) error {
	p.config = defaultConfig

	if len(config) > 0 {
		c := config[0]

		if c.FileExt != "" {
			p.config.FileExt = c.FileExt
		}

		if c.FileNamePrefix != "" {
			p.config.FileNamePrefix = c.FileNamePrefix
		}

		if c.FileNameSuffix != "" {
			p.config.FileNameSuffix = c.FileNameSuffix
		}

		if c.PastebinContentFileLocation != "" {
			p.config.PastebinContentFileLocation = c.PastebinContentFileLocation
		}

		if c.PastebinContentFolderName != "" {
			p.config.PastebinContentFolderName = c.PastebinContentFolderName
		}
	}

	err := os.MkdirAll(p.GetPastebinContentStoragePath(), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (p *pastebinContentLocalStorage) Save(content []byte) (string, error) {
	fileName := fmt.Sprintf("%s%s-%s%s%s", p.config.FileNamePrefix, common.GenerateRandomString(6), strconv.Itoa(int(time.Now().Unix())), p.config.FileNameSuffix, p.config.FileExt)

	file, err := os.Create(fmt.Sprintf("%s/%s", p.GetPastebinContentStoragePath(), fileName))
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func (p *pastebinContentLocalStorage) GetPastebinContentStoragePath() string {
	return fmt.Sprintf("%s/%s", p.config.PastebinContentFileLocation, p.config.PastebinContentFolderName)
}
