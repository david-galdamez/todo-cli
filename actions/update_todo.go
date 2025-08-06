package actions

import (
	"fmt"
	"log"

	"github.com/david-galdamez/todo-cli/database"
	"github.com/david-galdamez/todo-cli/utils"
)

func MarkCompleted(idArg string, dbConn *database.DBConnection) {
	todoId, err := utils.ParseUint(idArg)
	if err != nil {
		log.Fatalf("Error parsing the id: %v\n", err)
	}

	todo, err := dbConn.GetTodo(todoId)
	if err != nil {
		log.Fatalf("Error getting the todo: %v\n", err)
	}

	if todo.IsCompleted {
		fmt.Printf("Todo with id: %v is already completed\n", todoId)
		return
	}

	_, err = dbConn.MarkAsCompleted(todoId)
	if err != nil {
		log.Fatalf("Error marking as complete: %v\n", err)
	}

	fmt.Printf("Todo with id: %v mark as completed\n", todoId)
}

func UpdateTodo(args []string, dbConn *database.DBConnection) {
	todoId, err := utils.ParseUint(args[0])
	if err != nil {
		log.Fatalf("Error parsing the id: %v\n", err)
	}

	//args[1] is the new todo
	_, err = dbConn.UpdateTodo(todoId, args[1])
	if err != nil {
		log.Fatalf("Error updating todo: %v\n", err)
	}

	fmt.Printf("Todo with id: %v updated correctly\n", todoId)
}
