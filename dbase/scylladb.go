package dbase

import (
	"github.com/gocql/gocql"
	"gin-lnp-api/app"
	"time"
	"log"
)

type Database struct {
	*gocql.Session
}

func Connect() *gocql.Session {

	// connect to the database
	cluster := gocql.NewCluster(app.Config.CassHost)
	cluster.Keyspace = app.Config.Keyspace
	cluster.Consistency = gocql.LocalQuorum
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: app.Config.ClusterUsername,
		Password: app.Config.ClusterPassword,
	}

	// This is time after which the creation of session call would timeout.
	// This can be customised as needed.
	cluster.Timeout = app.Config.Timeout * time.Second
	cluster.ProtoVersion = app.Config.ProtoVersion

	// Pooling options
	dc := app.Config.DataCenter
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy(dc)
	cluster.HostFilter = gocql.DataCentreHostFilter(dc)

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Error creating Cassandra session: %v", err)
		panic(err)
	}

	return session
}


