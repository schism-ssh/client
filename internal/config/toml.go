package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

func LoadConfigFile(filename string) (conf *SchismConfig, err error) {
	conf = &SchismConfig{}
	_, err = toml.DecodeFile(filename, conf)
	return conf, err
}

func SaveConfigFile(filename string, conf *SchismConfig) error {
	confFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0660)
	if err != nil {
		return fmt.Errorf("%s\nerror opening %s", err, filename)
	}
	// This is fine to ignore the potential errors from
	// since we're later making a call to `Sync()`
	defer confFile.Close()

	tomlEnc := toml.NewEncoder(confFile)
	terr := tomlEnc.Encode(conf)
	if terr != nil {
		return fmt.Errorf("%s\nerror encoding TOML to file %s", terr, filename)
	}
	return confFile.Sync()
}
