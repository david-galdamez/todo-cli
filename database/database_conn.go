package database

import (
	"database/sql"
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

	return db.DB.Exec(`UPDATE todo SET is_completed = true, updated_at = ?  WHERE id = ?;`, updatedAt, todoId)
}

func (db *DBConnection) UpdateTodo(todoId uint, newTodo string) (sql.Result, error) {

	updatedAt := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	return db.DB.Exec("UPDATE todo SET todo = ?, updated_at = ? WHERE id = ?", newTodo, updatedAt, todoId)
}

func (db *DBConnection) GetTodo(todoId uint) (*models.Todo, error) {

	result, err := db.DB.Query("SELECT id, todo, is_completed, created_at, updated_at, due_to FROM todo WHERE id = ?;", todoId)
	if err != nil {
		return nil, err
	}

	todo := models.Todo{}

	err = result.Scan(todo.ID, todo.Todo, todo.IsCompleted, todo.CreatedAt, todo.UpdatedAt, todo.DueTo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (db *DBConnection) GetTodos() ([]models.Todo, error) {

	listOfTodos := []models.Todo{}

	result, err := db.DB.Query("SELECT id, todo, is_completed, created_at, updated_at, due_to FROM todo;")
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
