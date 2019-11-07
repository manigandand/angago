package main

import (
	"log"
	"os"
	"path/filepath"
)

func getConfigOrDefault() *AngaGoConf {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configFilePath, _ := filepath.Abs(home + "/.angago/config.yaml")
	isFileExist := fileExists(configFilePath)
	if !isFileExist {
		return loadFromDefault(home, configFilePath)
	}

	// load from the file
	body, err := getBytesForFileOrURL(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	config, err := decodeConfig(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("configuration loaded...")
	return config
}

func loadFromDefault(home, configFilePath string) *AngaGoConf {
	defaultConfigpath := "stub/default.config.yaml"
	body, err := getBytesForFileOrURL(defaultConfigpath)
	if err != nil {
		log.Fatal(err)
	}

	config, err := decodeConfig(body)
	if err != nil {
		log.Fatal(err)
	}
	cloneDefaultConfig(home, configFilePath, body)

	log.Println("Default configuration loaded. Please update your configuration at ", configFilePath)
	return config
}

func cloneDefaultConfig(home, configFilePath string, content []byte) {
	fileDir := home + "/.angago"
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		if err := os.Mkdir(fileDir, os.ModePerm); err != nil {
			log.Println("Can't create config dir ", fileDir, err)
			return
		}
	}

	f, err := os.Create(configFilePath)
	if err != nil {
		log.Println("Can't clone default config", err)
		return
	}
	if _, err := f.Write(content); err != nil {
		log.Println("Can't clone default config", err)
		f.Close()
		return
	}
	if err = f.Close(); err != nil {
		log.Println(err)
		return
	}
}
