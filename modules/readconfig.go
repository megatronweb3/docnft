package modules

import "os"
import "goland/config"
import "gopkg.in/yaml.v2"

func Config() *config.Config {
	f, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg config.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	return &cfg
}
