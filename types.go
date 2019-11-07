package main

import (
	"errors"
	"fmt"
	"sync"
)

var validResurcesMap = map[string]bool{
	"domains": true,
}

// AngaGoConf holds the reverse proxy configuration
type AngaGoConf struct {
	mu           sync.Mutex
	APIVersion   string `yaml:"apiVersion" json:"api_version"`
	APINamespace string `yaml:"apiNamespace" json:"api_namespace"`
	Kind         string `yaml:"kind" json:"kind"`
	Meta         struct {
		ShortName string `yaml:"shortName" json:"short_name"`
	} `yaml:"meta" json:"meta"`
	TLS     bool              `yaml:"tls" json:"tls"`
	Domains map[string]string `yaml:"domains" json:"domains"`
}

// TargetURL returns target proxy url
func (a *AngaGoConf) TargetURL(host string) (string, error) {
	a.mu.Lock()
	targetURL, ok := a.Domains[host]
	if !ok {
		a.mu.Unlock()
		return "", errors.New("invalid domain")
	}
	a.mu.Unlock()
	return targetURL, nil
}

// AbsTargetURL appends a valid scheme with targetURL based on configuration
func (a *AngaGoConf) AbsTargetURL(targetURL string) string {
	scheme := "http"
	if a.TLS {
		scheme = "https"
	}

	return fmt.Sprintf("%s://%s", scheme, targetURL)
}
