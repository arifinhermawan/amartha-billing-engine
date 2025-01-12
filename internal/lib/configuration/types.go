package configuration

type AppConfig struct {
	Database DatabaseConfig `yaml:"database"`
	Hash     HashKeyConfig  `yaml:"hash_key"`
}

type DatabaseConfig struct {
	Driver         string `yaml:"driver"`
	Host           string `yaml:"host"`
	DatabaseName   string `yaml:"database_name"`
	Password       string `yaml:"password"`
	Port           int    `yaml:"port"`
	Username       string `yaml:"username"`
	DefaultTimeout int    `yaml:"default_timeout"`
}

type HashKeyConfig struct {
	Password string `yaml:"password"`
}
