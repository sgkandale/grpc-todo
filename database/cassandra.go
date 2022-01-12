package database

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"path/filepath"
	"time"

	config "grpc-todo/config"

	"github.com/gocql/gocql"
)

var (
	Keyspace = config.Config.Cassandra.Keyspace
	Table    = config.Config.Cassandra.Table
)

func CassandraClientConn() *gocql.Session {
	log.Println("connecting Cassandra Read Client")

	cluster := gocql.NewCluster(config.Config.Cassandra.Host)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: config.Config.Cassandra.Username,
		Password: config.Config.Cassandra.Password,
	}

	certPath, err1 := filepath.Abs(config.Config.Cassandra.CertPath)
	keyPath, err2 := filepath.Abs(config.Config.Cassandra.KeyPath)
	caPath, err3 := filepath.Abs(config.Config.Cassandra.CaPath)

	if err1 != nil {
		log.Fatal("Unable to ReadFile('"+certPath+"') for Cassandra Client : ", err1)
	}
	if err2 != nil {
		log.Fatal("Unable to ReadFile('"+keyPath+"') for Cassandra Client : ", err2)
	}
	if err3 != nil {
		log.Fatal("Unable to ReadFile('"+config.Config.Cassandra.CaPath+"') for Cassandra Client : ", err3)
	}
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatal("Unable to LoadX509KeyPair() for Cassandra Client : ", err)
	}
	caCert, err := os.ReadFile(config.Config.Cassandra.CaPath)
	if err != nil {
		log.Println(err)
		log.Fatal("Unable to ReadFile('"+caPath+"') for Cassandra Client : ", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	cluster.SslOpts = &gocql.SslOptions{
		Config:                 tlsConfig,
		EnableHostVerification: false,
	}
	cluster.Hosts = []string{config.Config.Cassandra.Host + ":" + config.Config.Cassandra.Port}
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 20
	cluster.DisableInitialHostLookup = true
	cluster.Keyspace = Keyspace

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Error creating Cassandra client : ", err)
	}

	log.Println("Cassandra Client connected")

	return session
}

var CassandraClient = CassandraClientConn()
