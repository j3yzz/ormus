package main

import (
	"github.com/gocql/gocql"
	"github.com/ormushq/ormus/source/adapter/scylladb/initializedb"
	"log"
)

func main() {
	keyspace := "ormus_source2"
	hosts := []string{"127.0.0.1:9042"}
	connection := initializedb.NewScyllaDBConnection(gocql.Quorum, keyspace, hosts...)

	err := initializedb.CreateKeySpace(gocql.Quorum, keyspace, hosts...)
	if err != nil {
		log.Fatal(err)
	}

	session, err := initializedb.GetConnection(connection)
	if err != nil {
		log.Fatal("Failed to get ScyllaDB session:", err)
	}
	defer session.Close()

	err = initializedb.RunMigrations(connection, "./source/repository/scylladb/")
	if err != nil {
		log.Fatalln(err)
	}

}
