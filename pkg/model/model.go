package model

import "github.com/edgedb/edgedb-go"

type Person struct {
	edgedb.Optional

	ID        edgedb.UUID        `edgedb:"id"`
	FirstName string             `edgedb:"first_name"`
	LastName  edgedb.OptionalStr `edgedb:"last_name"`
	FullName  edgedb.OptionalStr `edgedb:"full_name"`
}

type Movie struct {
	ID       edgedb.UUID          `edgedb:"id"`
	Title    string               `edgedb:"title"`
	Year     edgedb.OptionalInt64 `edgedb:"year"`
	Director Person               `edgedb:"director"`
	Actors   []Person             `edgedb:"actors"`
}
