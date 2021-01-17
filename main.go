package main

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"github.com/vmkevv/duiztapi/ent"
)

func start() error {
	client, err := ent.Open("postgres", "postgres://postgres:12345@localhost:5432/duiztdb?sslmode=disable")
	if err != nil {
		return err
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		return err
	}
	return nil
}

func main() {
	err := parseJs()
	if err != nil {
		log.Fatalf("Error initalizing app: %v", err)
	}
}
