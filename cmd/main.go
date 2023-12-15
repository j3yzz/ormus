package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/ormushq/ormus/source/infrastructure/db/scylladb/initializedb"
	"log"
	"reflect"
)

func main() {
	cons := gocql.Quorum
	keyspace := "ormus_source"
	hosts := []string{"127.0.0.1:9042"}

	err := initializedb.CreateKeySpace(cons, keyspace, hosts...)
	if err != nil {
		log.Fatalln(err)
	}

	connection := initializedb.NewScyllaDBConnection(cons, keyspace, hosts...)
	session, err := initializedb.GetConnection(connection)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(reflect.TypeOf(session))

	err = initializedb.RunMigrations(connection, "./source/infrastructure/db/scylladb")
	if err != nil {
		log.Fatalln(err)
	}
	//
	//repository := scylladb_repo.NewEventRepository(session)
	//
	//_, err = repository.Save(model.Event{
	//	Id:        "1",
	//	Key:       "12sazx",
	//	Value:     "asdzxca",
	//	CreatedAt: time.Now(),
	//	UpdatedAt: time.Now(),
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}
}
