package config

import "github.com/spf13/viper"

type Config struct {
	PastebinConfig PastebinConfig `mapstructure:"PASTEBIN_CONFIG"`
}

type PastebinConfig struct {
	PastebinContentFileLocation string `mapstructure:"PASTEBIN_CONTENT_FILE_LOCATION"`
	PastebinContentFolderName   string `mapstructure:"PASTEBIN_CONTENT_FOLDER_NAME"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	c := Config{}

	err = viper.Unmarshal(&c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
