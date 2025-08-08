package actions

import (
	"flag"
	"fmt"
	"log"

	"github.com/david-galdamez/todo-cli/database"
)

func DeleteTodo(args []string, dbConn *database.DBConnection) {
	listCmd := flag.NewFlagSet("update", flag.ExitOnError)
	todoId := listCmd.Uint("id", 0, "Id of the todo you want to delete")

	listCmd.Parse(args)

	_, err := dbConn.DeleteTodo(*todoId)
	if err != nil {
		log.Fatalf("Error deleting the todo: %v\n", err)
	}

	fmt.Printf("Todo deleted correctly: %v\n", todoId)
}
