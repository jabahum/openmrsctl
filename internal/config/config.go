package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Deployment    string `yaml:"deployment"`
	OpenMRSHome   string `yaml:"openmrs_home"`
	TomcatService string `yaml:"tomcat_service"`
	MySQLService  string `yaml:"mysql_service"`
	BackupDir     string `yaml:"backup_dir"`
	LogsDir       string `yaml:"logs_dir"`
}

// SaveConfig writes config to .openmrsctl.yaml
func SaveConfig(c Config) error {
	data, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}
	return os.WriteFile(".openmrsctl.yaml", data, 0644)
}

// LoadConfig reads the config
func LoadConfig() (Config, error) {
	var c Config
	data, err := os.ReadFile(".openmrsctl.yaml")
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal(data, &c)
	return c, err
}
