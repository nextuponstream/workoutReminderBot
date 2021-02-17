package neo4j

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/nextuponstream/workoutReminderBot/pkg/entities"
)

func (n *Neo4j) InsertExercise(e entities.Exercise) error {
	item, err := n.insertItem()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(item)
	return nil
}

func (n *Neo4j) insertItem() (*Item, error) {
	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	// FIXME TransactionExecutionLimit: Timeout after 23 attempts,
	// last error: ConnectivityError: Unable to retrieve routing table from ngdb:7687:
	// dial tcp: lookup ngdb on 127.0.0.11:53: no such host
	result, err := session.WriteTransaction(createItemFn)
	if err != nil {
		log.Println("error writing transaction")
		return nil, err
	}
	return result.(*Item), nil
}

func createItemFn(tx neo4j.Transaction) (interface{}, error) {
	query := "CREATE (n:Item { id: $id, name: $name }) RETURN n.id, n.name"
	records, err := tx.Run(query, map[string]interface{}{
		"id":   1,
		"name": "Item 1",
	})
	// In face of driver native errors, make sure to return them directly.
	// Depending on the error, the driver may try to execute the function again.
	if err != nil {
		log.Print("error writing record")
		return nil, err
	}
	record, err := neo4j.Single(records, err)
	if err != nil {
		log.Print("error single")
		return nil, err
	}
	// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`
	return record, nil
}
