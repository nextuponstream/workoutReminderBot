package neo4j

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Item struct {
	Id   int64
	Name string
}

type Neo4j struct {
	driver neo4j.Driver
}

// Create a neo4j graph database
func Create(user string, password string) Neo4j {
	n := Neo4j{}
	// TODO connexion string
	driver, err := neo4j.NewDriver("neo4j://ngdb:7687", neo4j.BasicAuth(user, password, ""))
	if err != nil {
		log.Fatal(err)
	}

	n.driver = driver

	return n
}

func (n Neo4j) Close() {
	n.driver.Close()
}
