package config

type ServerConfig struct {
	Port       string
	TLSEnabled bool
	CertPath   string
	KeyPath    string
}

type DynamoDB struct {
	Region    string
	Table     string
	AccessKey string
	SecretKey string
}

type config struct {
	ServerConfig ServerConfig
	DynamoDB     DynamoDB
}
