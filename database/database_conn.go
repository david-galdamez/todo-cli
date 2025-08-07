package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/david-galdamez/todo-cli/models"
)

type DBConnection struct {
	DB *sql.DB
}

func (db *DBConnection) InsertTodo(todo models.Todo) (sql.Result, error) {

	return db.DB.Exec(` INSERT INTO todo (todo, is_completed, created_at, updated_at, due_to)
	VALUES (?, ?, ?, ?, ?);`, todo.Todo, todo.IsCompleted, todo.CreatedAt, todo.UpdatedAt, todo.DueTo)
}

func (db *DBConnection) DeleteTodo(todoId uint) (sql.Result, error) {
	return db.DB.Exec(`DELETE FROM todo WHERE id = ?;`, todoId)
}

func (db *DBConnection) MarkAsCompleted(todoId uint) (sql.Result, error) {

	updatedAt := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	dueTo := sql.NullTime{
		Valid: false,
	}

	return db.DB.Exec(`UPDATE todo SET is_completed = true, updated_at = ?, due_to = ?  WHERE id = ?;`, updatedAt, dueTo, todoId)
}

func (db *DBConnection) UpdateTodo(todoId *uint, newTodo *string, newDate *time.Time) (sql.Result, error) {

	fields := []string{}
	args := []interface{}{}

	if *newTodo != "" {
		fields = append(fields, "todo = ?")
		args = append(args, *newTodo)
	}

	if newDate != nil {

		fields = append(fields, "due_to = ?")
		args = append(args, sql.NullTime{
			Time:  *newDate,
			Valid: true,
		})
	}

	updatedAt := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	fields = append(fields, "updated_at = ?")
	args = append(args, updatedAt)

	args = append(args, *todoId)

	queryString := fmt.Sprintf("UPDATE todo SET %v WHERE id = ?", joinFields(fields))

	return db.DB.Exec(queryString, args...)
}

func joinFields(fields []string) string {
	return strings.Join(fields, ",")
}

func (db *DBConnection) GetTodo(todoId uint) (*models.Todo, error) {

	row := db.DB.QueryRow("SELECT id, todo, is_completed, created_at, updated_at, due_to FROM todo WHERE id = ?;", todoId)

	todo := models.Todo{}

	err := row.Scan(&todo.ID, &todo.Todo, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt, &todo.DueTo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (db *DBConnection) GetTodos(completed *bool) ([]models.Todo, error) {

	listOfTodos := []models.Todo{}
	queryString := "SELECT id, todo, is_completed, created_at, updated_at, due_to FROM todo"

	if *completed {
		queryString = queryString + " WHERE is_completed = true;"
	} else {
		queryString = queryString + ";"
	}

	result, err := db.DB.Query(queryString)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		todo := models.Todo{}

		err := result.Scan(&todo.ID, &todo.Todo, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt, &todo.DueTo)
		if err != nil {
			result.Close()
			return nil, err
		}

		listOfTodos = append(listOfTodos, todo)
	}

	return listOfTodos, nil
}
