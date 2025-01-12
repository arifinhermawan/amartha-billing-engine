package configuration

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func (c *Configuration) GetConfig() *AppConfig {
	c.doLoadConfigOnce.Do(func() {
		cfg, err := c.loadConfig()
		if err != nil {
			log.Fatalf("[GetConfig] c.loadConfig() got error: %v\n", err)
		}

		c.config = cfg
	})

	return &c.config
}

func (c *Configuration) loadConfig() (AppConfig, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	filePath := fmt.Sprintf("etc/files/config.%s.yaml", env)
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("[loadConfig] os.ReadFile() got error: %v\n", err)
		return AppConfig{}, err
	}

	var config AppConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("[loadConfig] yaml.Unmarshal() got error: %v\n", err)
		return AppConfig{}, err
	}

	return config, nil
}
