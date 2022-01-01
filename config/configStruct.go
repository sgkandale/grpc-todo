package config

type ServerConfig struct {
	Port       string
	TLSEnabled bool
	CertPath   string
	KeyPath    string
}

type config struct {
	ServerConfig ServerConfig
}
