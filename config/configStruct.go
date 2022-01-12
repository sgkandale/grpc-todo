package config

type ServerConfig struct {
	Port       string
	TLSEnabled bool
	CertPath   string
	KeyPath    string
}

type Cassandra struct {
	Host     string
	Port     string
	Username string
	Password string
	Keyspace string
	Table    string
	CertPath string
	KeyPath  string
	CaPath   string
}

type config struct {
	ServerConfig ServerConfig
	Cassandra    Cassandra
}
