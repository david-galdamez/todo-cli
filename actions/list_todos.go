package actions

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/david-galdamez/todo-cli/database"
)

func ListTodos(args []string, dbConn *database.DBConnection) {

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	completed := listCmd.Bool("completed", false, "When true, filters the todo list showing only the ones that are complete")

	listCmd.Parse(args)

	todos, err := dbConn.GetTodos(completed)
	if err != nil {
		log.Fatalf("Error getting the todos: %v\n", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 3, 3, 4, ' ', 0)
	_, err = fmt.Fprintf(w, "ID\tTodo\tCompleted\tCreated At\tUpdated At\tDue To\t\n")
	if err != nil {
		log.Fatalf("Error printing table: %v\n", err)
	}

	if len(todos) > 0 {
		for _, todo := range todos {

			updatedAt := "-"
			dueTo := "-"

			if todo.UpdatedAt.Valid {
				updatedAt = todo.UpdatedAt.Time.Format(time.DateTime)
			}

			if todo.DueTo.Valid {
				dueTo = todo.DueTo.Time.Format(time.DateTime)
			}

			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t\n", todo.ID, todo.Todo, todo.IsCompleted, todo.CreatedAt.Format(time.DateTime), updatedAt, dueTo)
		}
	}

	w.Flush()
}
