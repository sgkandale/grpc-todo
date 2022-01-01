package config

import "log"

func VerifyConfig() config {

	if parsedConfig.ServerConfig.Port == "" {
		log.Fatal("No port specified")
	}

	if parsedConfig.ServerConfig.TLSEnabled {
		if parsedConfig.ServerConfig.CertPath == "" {
			log.Fatal("no CertPath provided")
		}
		if parsedConfig.ServerConfig.KeyPath == "" {
			log.Fatal("no KeyPath provided")
		}
	}

	return parsedConfig

}

var Config = VerifyConfig()
