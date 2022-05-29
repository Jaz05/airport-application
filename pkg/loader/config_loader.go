package loader

import (
	"airport/internal/path"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfig[T any](fileName string, cfg *T) {
	var configPath = path.GetRootPath() + "/config/" + fileName
	f, err := os.Open(configPath)
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
