package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func getBytesForFileOrURL(path string) ([]byte, error) {
	u, err := url.ParseRequestURI(path)
	if err == nil && u.Scheme != "" {
		return getBytesFromURL(path)
	}

	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	return ioutil.ReadFile(path)
}

func getBytesFromURL(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func decodeConfig(in []byte) (*AngaGoConf, error) {
	var config AngaGoConf
	if err := yaml.Unmarshal(in, &config); err != nil {
		return nil, err
	}

	// TODO: validate other basic configs
	if _, ok := validResurcesMap[strings.ToLower(config.Kind)]; !ok {
		return nil, errors.New("invalid resource type")
	}

	return &config, nil
}
