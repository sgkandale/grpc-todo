package config

type ServerConfig struct {
	Port       string
	TLSEnabled bool
	CertPath   string
	KeyPath    string
}

type MySQL struct {
	Username string
	Password string
	Host     string
	DBName   string
}

type config struct {
	ServerConfig ServerConfig
	MySQL        MySQL
}
