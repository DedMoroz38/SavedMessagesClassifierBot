package main

import (
	"context"
	"log"
	"saved_messages_classifier/classifier"

	"github.com/jackc/pgx/v5"
)

func NewDBService() *classifier.Queries {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=postgres dbname=tg_classifier password=postgres host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Println("Error while connecting to db: ", err)
	}
	defer conn.Close(ctx)

	queries := classifier.New(conn)

	return queries
}
