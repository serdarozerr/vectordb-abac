package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type ErrConfig struct {
	Msg  string
	Code int
}

func (e ErrConfig) Error() string {
	return e.Msg
}

var (
	NotFoundErr = ErrConfig{Msg: "not found", Code: 404}
	FileError   = ErrConfig{Msg: "file read error", Code: 500}
	InvalidErr  = ErrConfig{Msg: "invalid parameter", Code: 400}
	EmptyValue  = ErrConfig{Msg: "empty value", Code: 400}
	YamlInvalid = ErrConfig{Msg: "invalid Yaml", Code: 400}
)

type Config struct {
	App struct {
		Name    string `yaml:"Name"`
		Version string `yaml:"Version"`
		Port    int    `yaml:"Port"`
	} `yaml:"App"`
	VectorDB struct {
		Type   string `yaml:"Type"`
		URL    string `yaml:"URL"`
		Port   int    `yaml:"Port"`
		ApiKey string `yaml:"ApiKey"`
	} `yaml:"VectorDB"`

	LLM struct {
		Type            string `yaml:"Type"`
		ApiKey          string `yaml:"ApiKey"`
		EmbeddingModel  string `yaml:"EmbeddingModel"`
		CompletionModel string `yaml:"CompletionModel"`
		VectorDimension int    `yaml:"VectorDimension"`
	} `yaml:"LLM"`
}

type GetEnvValue interface {
	GetValue([]string) (string, error)
}

func (cfg *Config) YamlUnmarshal(data []byte) error {

	if len(data) == 0 {
		return ErrConfig{Msg: "YAML data is empty", Code: 400}
	}

	// Unmarshal YAML
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return ErrConfig{Msg: "Failed to parse YAML: " + err.Error(), Code: 400}
	}

	// Validate fields
	if cfg.App.Name == "" {
		return ErrConfig{Msg: "AppName is missing", Code: 400}
	}
	if cfg.App.Version == "" {
		return ErrConfig{Msg: "AppVersion is missing", Code: 400}
	}
	if cfg.VectorDB.URL == "" {
		return ErrConfig{Msg: "Vector DB URL is missing", Code: 400}
	}
	if cfg.VectorDB.Port > 66000 || cfg.VectorDB.Port < 1000 {
		return ErrConfig{Msg: fmt.Sprintf("QdrantPort %d is out of range (1000-66000)", cfg.VectorDB.Port), Code: 400}
	}

	if cfg.App.Port > 66000 || cfg.App.Port < 1000 {
		return ErrConfig{Msg: fmt.Sprintf("App port %d is out of range (1000-66000)", cfg.App.Port), Code: 400}
	}

	return nil

}
func ReadYamlFile(filePath string) (*Config, error) {
	if filePath == "" {
		InvalidErr.Msg = "File path is missing"
		return nil, InvalidErr
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, FileError
	}
	conf := &Config{}
	err = conf.YamlUnmarshal(data)
	if err != nil {
		return nil, err
	}

	return conf, nil

}
