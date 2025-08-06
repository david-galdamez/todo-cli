package actions

import (
	"fmt"
	"log"

	"github.com/david-galdamez/todo-cli/database"
	"github.com/david-galdamez/todo-cli/utils"
)

func DeleteTodo(idArg string, dbConn *database.DBConnection) {
	todoId, err := utils.ParseUint(idArg)
	if err != nil {
		log.Fatalf("Error parsing id: %v\n", err)
	}

	_, err = dbConn.DeleteTodo(todoId)
	if err != nil {
		log.Fatalf("Error deleting the todo: %v\n", err)
	}

	fmt.Printf("Todo deleted correctly: %v\n", todoId)
}
