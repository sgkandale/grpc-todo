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

	if parsedConfig.Cassandra.Host == "" {
		log.Fatal("No Cassandra Host specified")
	}
	if parsedConfig.Cassandra.Port == "" {
		log.Fatal("No Cassandra Port specified")
	}
	if parsedConfig.Cassandra.Username == "" {
		log.Fatal("No Cassandra Username specified")
	}
	if parsedConfig.Cassandra.Password == "" {
		log.Fatal("No Cassandra Password specified")
	}
	if parsedConfig.Cassandra.Keyspace == "" {
		log.Fatal("No Cassandra Keyspace specified")
	}
	if parsedConfig.Cassandra.Table == "" {
		log.Fatal("No Cassandra Table specified")
	}
	if parsedConfig.Cassandra.CertPath == "" {
		log.Fatal("No Cassandra CertPath specified")
	}
	if parsedConfig.Cassandra.KeyPath == "" {
		log.Fatal("No Cassandra KeyPath specified")
	}
	if parsedConfig.Cassandra.CaPath == "" {
		log.Fatal("No Cassandra CaPath specified")
	}

	return parsedConfig

}

var Config = VerifyConfig()
