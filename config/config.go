package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/olebedev/config"
)

const ConfigName = "config_client.yaml"

var configpath string

func init() {

	root, _ := os.Getwd()

	configpath = filepath.Join(root, ConfigName)
	/*

		if _, err := os.Stat(configpath); os.IsNotExist(err) {
			fmt.Printf("no such file or directory : %s\n", configpath)
			panic(err)
		}
	*/

}

func GetConfig(configfile string) *config.Config {

	if configfile != "" {
		configpath = configfile
	}

	cfg, err := config.ParseYamlFile(configpath)

	if err != nil {
		fmt.Errorf("config file parser error %v\n", err)
	}

	return cfg
}
