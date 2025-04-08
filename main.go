package main

import (
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// define the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Port = 9042
	cluster.Keyspace = "ai_parking"

	// create a session
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Scylladb connection established")
	}
	defer session.Close()
}
