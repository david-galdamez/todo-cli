package actions

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/david-galdamez/todo-cli/database"
	"github.com/david-galdamez/todo-cli/models"
)

func AddToDo(args []string, dbConn *database.DBConnection) {

	newTodo := models.Todo{
		ID:        0,
		Todo:      args[0],
		CreatedAt: time.Now(),
		UpdatedAt: sql.NullTime{Valid: false},
		DueTo:     sql.NullTime{Valid: false},
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
