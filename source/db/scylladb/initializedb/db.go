package initializedb

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/ormushq/ormus/source/db/scylladb"
	"log"
	"time"
)

type scyllaDBConnection struct {
	consistency gocql.Consistency
	keyspace    string
	hosts       []string
}

func (conn *scyllaDBConnection) createCluster() *gocql.ClusterConfig {
	cluster := gocql.NewCluster(conn.hosts...)
	cluster.Consistency = conn.consistency
	cluster.Keyspace = conn.keyspace
	cluster.Timeout = 5 * time.Second
	cluster.RetryPolicy = &gocql.ExponentialBackoffRetryPolicy{
		NumRetries: 5,
		Min:        time.Second,
		Max:        10 * time.Second,
	}
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	log.Println("cluster was created.")
	return cluster
}

func (conn *scyllaDBConnection) createSession(cluster *gocql.ClusterConfig) (scylladb.SessionxInterface, error) {
	session, err := scylladb.WrapSession(cluster.CreateSession())
	if err != nil {
		fmt.Println("an error occureed while creating DB Session", err.Error())
		return nil, err
	}
	log.Println("session was created")
	return session, nil
}
