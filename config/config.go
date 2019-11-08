package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/manigandand/angago/util"
)

// Env ...
var (
	Env                     string
	Port                    string
	Home                    string
	AngagoConfigPath        string
	DefaultAngagoConfigPath string
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	Home = home
	DefaultAngagoConfigPath = home + "/.angago/config.yaml"

	GetAllEnv()
}

// GetAllEnv should get all the env configs required for the app.
func GetAllEnv() {
	// API Configs
	mustEnv("ENV", &Env, "dev")
	mustEnv("PORT", &Port, "80")
	mustEnv("ANGAGO_CONFIG_PATH", &AngagoConfigPath, DefaultAngagoConfigPath)

	if !filepath.IsAbs(AngagoConfigPath) || !util.FileExists(AngagoConfigPath) {
		log.Printf("the given ANGAGO_CONFIG_PATH=%s is not absolute path or file not found, using default value %s\n",
			AngagoConfigPath, DefaultAngagoConfigPath)
	}
}

// mustEnv get the env variable with the name 'key' and store it in 'value'
func mustEnv(key string, value *string, defaultVal string) {
	if *value = os.Getenv(key); *value == "" {
		*value = defaultVal
		log.Printf("%s env variable not set, using default value %s.\n",
			key, defaultVal)
	}
}
