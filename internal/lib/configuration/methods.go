package configuration

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// GetConfig retrieves the application configuration in a thread-safe manner.
// It applies singleton pattern so it ensures that the configuration is loaded only once,
// even if the function is called multiple times. The configuration is loaded using the loadConfig()
// function, which loads it from a YAML file based on the environment setting.
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

// loadConfig loads the application configuration from a YAML file based on the
// current environment, defaulting to "development" if the variable is not set.
// The function reads the YAML file, unmarshals it into an AppConfig struct, and returns
// the configuration data.
//
// If any error occurs while reading the file or unmarshaling the YAML data,
// the function logs the error and returns an empty AppConfig struct.
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
