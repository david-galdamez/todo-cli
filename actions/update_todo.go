package actions

import (
	"flag"
	"fmt"
	"log"
	"time"

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

	listCmd := flag.NewFlagSet("update", flag.ExitOnError)

	todoId := listCmd.Uint("id", 0, "Id of the todo you want to update")
	newTodo := listCmd.String("title", "", "Title of the todo you want to update")
	newDate := listCmd.String("dueTo", "", "When the todo is due to")

	listCmd.Parse(args)

	var dueTo *time.Time

	if *newDate != "" {
		newTime, err := time.Parse("2006-01-02", *newDate)
		if err != nil {
			log.Fatalf("Error parsing date: %v\n", err)
		}

		dueTo = &newTime
	}

	if *newTodo == "" || dueTo == nil {
		fmt.Println("You have to pass a new todo title or a new due to date")
		return
	}

	_, err := dbConn.UpdateTodo(todoId, newTodo, dueTo)
	if err != nil {
		log.Fatalf("Error updating todo: %v\n", err)
	}

	fmt.Printf("Todo with id: %v updated correctly\n", *todoId)
}
