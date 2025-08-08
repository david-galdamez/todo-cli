package actions

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func HelpTodo() {
	w := tabwriter.NewWriter(os.Stdout, 3, 3, 4, ' ', 0)
	_, err := fmt.Fprintf(w, "Command\tDescription\n")
	if err != nil {
		log.Fatalf("Error printing manual: %v\n", err)
	}

	fmt.Fprintln(w)
	fmt.Fprint(w, "migrate\tMigrate the database for todos\n")
	fmt.Fprint(w, "drop\tDrop database of todos\n")
	fmt.Fprint(w, "list\tPrints a list of the registered todos\n")
	fmt.Fprint(w, "add\tAdd a todo to the database\n")
	fmt.Fprint(w, "update\tUpdated a specific todo\n")
	fmt.Fprint(w, "complete\tMarks a specific todo as completed\n")
	fmt.Fprint(w, "delete\tDeletes the specific todo\n")

	w.Flush()
}
