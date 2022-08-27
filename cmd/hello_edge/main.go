package main

import (
	"context"
	"fmt"
	"log"

	"hello_edge/pkg/model"

	"github.com/edgedb/edgedb-go"
)

func main() {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var actors []model.Person
	args := map[string]interface{}{"first_name": "John"}
	query := `SELECT Person {
		id,	first_name, last_name, full_name
	} FILTER .first_name = <str>$first_name
`
	err = client.Query(ctx, query, &actors, args)
	if err != nil {
		log.Fatal(err)
	}

	for i, actor := range actors {
		fmt.Println(i, ">>", actor)
	}

	var movies []model.Movie
	query = `SELECT Movie {
		id, title, year,
		director: {
			id,	first_name, last_name, full_name
		},
		actors: {
			id,	first_name, last_name, full_name
		}
	}`

	err = client.Query(ctx, query, &movies)
	if err != nil {
		log.Fatal(err)
	}

	for i, movie := range movies {
		fmt.Println(i, ">>", movie)
	}
}
