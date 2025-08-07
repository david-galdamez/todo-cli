package actions

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/david-galdamez/todo-cli/database"
	"github.com/david-galdamez/todo-cli/models"
)

func AddToDo(args []string, dbConn *database.DBConnection) {

	listCmd := flag.NewFlagSet("add", flag.ExitOnError)
	todoTitle := listCmd.String("todo", "", "Todo yo want to add")
	dueToDate := listCmd.String("dueTo", "", "When the todo is due to")
	listCmd.Parse(args)

	newTodo := models.Todo{
		ID:        0,
		Todo:      *todoTitle,
		CreatedAt: time.Now(),
		UpdatedAt: sql.NullTime{Valid: false},
		DueTo:     sql.NullTime{Valid: false},
	}

	if *dueToDate != "" {
		time, err := time.Parse("2006-01-02", *dueToDate)
		if err != nil {
			log.Fatalf("Error parsing time: %v\n", err)
		}

		newTodo.DueTo.Valid = true
		newTodo.DueTo.Time = time
	}

	result, err := dbConn.InsertTodo(newTodo)
	if err != nil {
		log.Fatalf("Error creating todo: %v\n", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Error getting the todo id: %v\n", err)
	}

	fmt.Printf("Todo created correctly: %v\n", id)
}
