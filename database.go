package main

import (
	"context"
	"fmt"
	"os"
	"log"

	"github.com/jackc/pgx"
)

func InitialiseFile() File {
	return File{FileName: "", Hash: ""}
}

func GenerateFileHash() string {
	return "random"
}

func ConnectToDatabase() {
	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	log.Fatalf("Config: %v", Conf.Env)

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:password@localhost:5432/files")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var id int64
	//var weight int64
	//err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	err = conn.QueryRow(context.Background(), "select * from files").Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("id", id)
}
