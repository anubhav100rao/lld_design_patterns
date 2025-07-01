package adapter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Target interface
type ConfigLoader interface {
	Load(path string, out interface{}) error
}

// Adaptee #1: JSON loader
type JSONLoader struct{}

func (j *JSONLoader) Decode(data []byte, out interface{}) error {
	return json.Unmarshal(data, out)
}

// Adaptee #2: YAML loader
type YAMLLoader struct{}

func (y *YAMLLoader) Decode(data []byte, out interface{}) error {
	return yaml.Unmarshal(data, out)
}

// Adapter – implements ConfigLoader
type FileConfigAdapter struct {
	jsonLoader *JSONLoader
	yamlLoader *YAMLLoader
}

func NewFileConfigAdapter() *FileConfigAdapter {
	return &FileConfigAdapter{
		jsonLoader: &JSONLoader{},
		yamlLoader: &YAMLLoader{},
	}
}

func (a *FileConfigAdapter) Load(path string, out interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	switch ext := path[len(path)-5:]; ext {
	case ".json":
		return a.jsonLoader.Decode(data, out)
	case ".yaml", ".yml":
		return a.yamlLoader.Decode(data, out)
	default:
		return fmt.Errorf("unsupported config format: %s", path)
	}
}

type AppConfig struct {
	Port int `json:"port" yaml:"port"`
}

func RunConfigLoaderDemo() {
	loader := NewFileConfigAdapter()
	var cfg AppConfig
	if err := loader.Load("config.yaml", &cfg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Loaded config: %+v\n", cfg)
}
