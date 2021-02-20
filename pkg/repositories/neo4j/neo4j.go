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

	// create constraints
	telegramUserConstraint := "CREATE CONSTRAINT telegramUser IF NOT EXISTS\n" +
		"ON (u:User)\n" +
		"ASSERT u.tid IS UNIQUE\n"
	activityConstraint := "CREATE CONSTRAINT activity IF NOT EXISTS\n" +
		"ON (a:Activity)\n" +
		"ASSERT a.name IS UNIQUE\n"

	queries := []string{telegramUserConstraint, activityConstraint}
	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	for _, q := range queries {
		_, err := session.Run(q, map[string]interface{}{})
		if err != nil {
			log.Fatal(err)
		}
	}

	return n
}

func (n Neo4j) Close() {
	n.driver.Close()
}
