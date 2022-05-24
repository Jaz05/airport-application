package loader

import (
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfig[T any](path string, cfg *T) {
	f, err := os.Open(path)
	if err != nil {
		println(err.Error())
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		println(err.Error())
	}
}
