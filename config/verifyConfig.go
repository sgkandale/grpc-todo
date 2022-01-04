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

	if parsedConfig.MySQL.Username == "" {
		log.Fatal("No MySQL username provided")
	}
	if parsedConfig.MySQL.Password == "" {
		log.Fatal("No MySQL password provided")
	}
	if parsedConfig.MySQL.Host == "" {
		log.Fatal("No MySQL host provided")
	}
	if parsedConfig.MySQL.DBName == "" {
		log.Fatal("No MySQL database name provided")
	}

	return parsedConfig

}

var Config = VerifyConfig()
