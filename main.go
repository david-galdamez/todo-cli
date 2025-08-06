package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/david-galdamez/todo-cli/actions"
	"github.com/david-galdamez/todo-cli/database"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	var action string

	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	conn, err := sql.Open("sqlite3", "./database/todo-database.db")
	if err != nil {
		log.Fatalf("Unexpected error %v\n", err)
	}

	defer conn.Close()

	dbConn := &database.DBConnection{
		DB: conn,
	}

	_, err = database.MigrateDatabase(dbConn)
	if err != nil {
		log.Fatalf("Error creating table: %v\n", err)
	}

	switch action {
	case "migrate":
		_, err = database.MigrateDatabase(dbConn)
		if err != nil {
			log.Fatalf("Error creating table: %v\n", err)
		}

		fmt.Print("Table migrated correctly")
	case "drop":
		_, err = database.DropDatabase(dbConn)
		if err != nil {
			log.Fatalf("Error droping table: %v\n", err)
		}

		fmt.Print("Table droped correctly")
	case "list":
		actions.ListTodos(os.Args[2:], dbConn)
	case "add":
		actions.AddToDo(os.Args[2:], dbConn)
	case "update":
		actions.UpdateTodo(os.Args[2:], dbConn)
	case "complete":
		actions.MarkCompleted(os.Args[2], dbConn)
	case "delete":
		actions.DeleteTodo(os.Args[2], dbConn)
	case "help":
		fmt.Printf("Todo actions %v\n", action)
	default:
		fmt.Println("Action does not exist")
	}
}
