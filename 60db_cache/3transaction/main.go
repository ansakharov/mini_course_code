package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func main() {
	// canceling of ctx closes connect
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres password=changeme sslmode=disable")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(time.Minute * 5)

	// squirrel
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
			id BIGSERIAL PRIMARY KEY,
			name TEXT
		)`)
	if err != nil {
		panic(err)
	}

	values := []interface{}{"Alice", "Bob", "Charlie"}
	valueStrings := make([]string, 0, len(values))
	valueArgs := make([]interface{}, 0, len(values))
	for idx, value := range values {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d)", idx+1))
		valueArgs = append(valueArgs, value)
	}

	query := fmt.Sprintf("INSERT INTO users (name) VALUES %s", strings.Join(valueStrings, ", "))
	_, err = db.ExecContext(context.Background(), query, valueArgs...)
	if err != nil {
		panic(err)
	}

	fmt.Println("users_before_delete")
	Users(ctx, db)

	// Select users for update and then delete them
	SelectForUpdateAndDelete(ctx, db)

	fmt.Println("users_after_delete")
	Users(ctx, db)
}

func Users(ctx context.Context, db *sqlx.DB) {
	rows, err := db.QueryContext(ctx, `SELECT * FROM users`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	fmt.Println(users)
}

func SelectForUpdateAndDelete(ctx context.Context, db *sqlx.DB) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	rows, err := tx.QueryxContext(ctx, `SELECT * FROM users FOR UPDATE`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	for _, user := range users {
		_, err = tx.ExecContext(ctx, `DELETE FROM users WHERE id = $1`, user.ID)
		if err != nil {
			panic(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
