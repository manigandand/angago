package main

// AngaGoConf holds the reverse proxy configuration
type AngaGoConf struct {
	APIVersion   string `yaml:"apiVersion" json:"api_version"`
	APINamespace string `yaml:"apiNamespace" json:"api_namespace"`
	Kind         string `yaml:"kind" json:"kind"`
	Meta         struct {
		ShortName string `yaml:"shortName" json:"short_name"`
	} `yaml:"meta" json:"meta"`
	TLS     bool              `yaml:"tls" json:"tls"`
	Domains map[string]string `yaml:"domains" json:"domains"`
}

var validResurcesMap = map[string]bool{
	"domains": true,
}

// DomainRegistryMap ...
var DomainRegistryMap = map[string]string{
	"app.gopherhut.com":    "http://localhost:8081",
	"status.gopherhut.com": "http://localhost:8082",
	"admin.gopherhut.com":  "http://localhost:8083",
}
