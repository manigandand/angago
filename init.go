package main

import (
	"log"
	"os"

	"github.com/manigandand/angago/config"
	"github.com/manigandand/angago/util"
)

func getConfigOrDefault() *AngaGoConf {
	isFileExist := util.FileExists(config.AngagoConfigPath)
	if !isFileExist {
		return loadFromDefault(config.Home, config.AngagoConfigPath)
	}

	// load from the file
	body, err := getBytesForFileOrURL(config.AngagoConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	angagoconfig, err := decodeConfig(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("configuration loaded from %s...\n", config.AngagoConfigPath)
	return angagoconfig
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
