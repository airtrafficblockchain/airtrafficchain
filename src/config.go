package main

import (
	"os"
)

type Config struct {
	switchName string
	switchHost string
	switchPort string
	senzieMode string
	senzieName string
	dotKeys    string
	idRsa      string
	idRsaPub   string
}

type CassandraConfig struct {
	host        string
	port        string
	keyspace    string
	consistancy string
}

type KafkaConfig struct {
	topic  string
	cgroup string
	khost  string
	kport  string
	zhost  string
	zport  string
}

type ApiConfig struct {
	poApi   string
	prntApi string
}

var config = Config{
	switchName: getEnv("SWITCH_NAME", "senzswitch"),
	switchHost: getEnv("SWITCH_HOST", "www.rahasak.com"),
	switchPort: getEnv("SWITCH_PORT", "7070"),
	senzieMode: getEnv("SENZIE_MODE", "dev"),
	senzieName: getEnv("SENZIE_NAME", "chain"),
	dotKeys:    getEnv("DOT_KEYS", ".keys"),
	idRsa:      getEnv("ID_RSA", ".keys/id_rsa"),
	idRsaPub:   getEnv("ID_RSA_PUB", ".keys/id_rsa.pub"),
}

var cassandraConfig = CassandraConfig{
	host:        getEnv("CASSANDRA_HOST", "dev.localhost"),
	port:        getEnv("CASSANDRA_PORT", "9042"),
	keyspace:    getEnv("CASSANDRA_KEYSPACE", "zchain"),
	consistancy: getEnv("CASSANDRA_CONSISTANCY", "ALL"),
}

var kafkaConfig = KafkaConfig{
	topic:  getEnv("KAFKA_TOPIC", "chain"),
	cgroup: getEnv("KAFKA_CGROUP", "chaing"),
	khost:  getEnv("KAFKA_KHOST", "dev.localhost"),
	kport:  getEnv("KAFKA_KPORT", "9092"),
	zhost:  getEnv("KAFKA_ZHOST", "dev.localhost"),
	zport:  getEnv("KAFKA_ZPORT", "2181"),
}

var apiConfig = ApiConfig{
	poApi:   getEnv("PO_API", "https://dev.localhost/v1.0/PurchaseOrder"),
	prntApi: getEnv("PRNT_API", "http://dev.localhost/v1.0/Print"),
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
