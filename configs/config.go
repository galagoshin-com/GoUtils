package configs

import (
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/files"
	"os"
)

type Config struct {
	Name    string
	version uint
	file    *files.File
	data    map[string]any
}

func (config *Config) Init(defaultData map[string]any, actualVersion uint) {
	filename := config.Name + ".gconf"
	config.file = &files.File{Path: filename}
	if !config.file.Exists() {
		logger.Print(fmt.Sprintf("Generating new \"%s\" configuration file.", filename))
		err := config.file.Create()
		if err != nil {
			logger.Error(err)
			return
		}
		err = config.file.WriteString(fmt.Sprintf("#v=%d\n%s", actualVersion, buildStr(defaultData)))
		if err != nil {
			logger.Error(err)
			return
		}
		err = config.file.Close()
		if err != nil {
			logger.Error(err)
			return
		}
		config.data = defaultData
	} else {
		err := config.file.Open(os.O_RDWR)
		if err != nil {
			logger.Error(err)
			return
		}
		content := config.file.ReadString()
		version, data := buildMap(content)
		if version != actualVersion {
			logger.Print(fmt.Sprintf("New version of \"%s\" config found.", config.Name))
			for key, _ := range data {
				ex := false
				for key1, _ := range defaultData {
					if key == key1 {
						ex = true
					}
				}
				if !ex {
					delete(data, key)
				}
			}
			for key, val := range defaultData {
				ex := false
				for key1, _ := range data {
					if key == key1 {
						ex = true
					}
				}
				if !ex {
					data[key] = val
				}
			}
			config.version = actualVersion
			config.data = data
			err = config.Save()
			if err != nil {
				logger.Error(err)
				return
			}
		} else {
			config.data = data
			config.version = actualVersion
		}
	}
	logger.Debug(2, false, fmt.Sprintf("Config name: %s", config.Name))
	logger.Debug(2, false, fmt.Sprintf("Config version: %d", config.version))
	logger.Debug(2, false, fmt.Sprintf("Config data: %+v", config.data))
}

func (config *Config) Save() error {
	err := config.file.Create()
	if err != nil {
		return err
	}
	err = config.file.WriteString(fmt.Sprintf("#v=%d\n%s", config.version, buildStr(config.data)))
	if err != nil {
		return err
	}
	err = config.file.Close()
	if err != nil {
		return err
	}
	return nil
}

func (cfg *Config) Get(key string) (any, bool) {
	res, found := cfg.data[key]
	return res, found
}

func (cfg *Config) Exists(key string) bool {
	_, found := cfg.data[key]
	return found
}

func (cfg *Config) Set(key string, val any) {
	cfg.data[key] = val
}
