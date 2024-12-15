package services

import (
	"encoding/json"
	"github.com/easy-lamb/easy-lamb-go/models"
	"github.com/sirupsen/logrus"
	"os"
)

func GetConfig(configPath string) (models.Config, error) {
	logrus.Debugf("Get configuration %s", configPath)

	configContent, err := os.ReadFile(configPath)

	if err != nil {
		return models.Config{}, err
	}

	var config models.Config

	err = json.Unmarshal(configContent, &config)

	return config, err
}
