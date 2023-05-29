package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres password=changeme sslmode=disable")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	//`SetMaxOpenConns`   позволяет настроить пул соединений, каждая горутина обращая к БД получит другой коннект. Всего может обрабатываться 15 соединений одновременно, последующие встанут в очередь.
	//`SetMaxIdleConns`  отвечает за кол-во простаивающих соединений. Если нагрузка долгое время слишком мала, то имеет смысл закрыть лишние соединения и данная настройка позволяет конфигурировать кол-во простойщиков.
	// `SetConnMaxIdleTime`  в свою очередь является таймаутом после которого неиспользуемое соединение будет закрыто.
	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(15)
	db.SetConnMaxIdleTime(time.Hour * 5)
	// db.SetConnMaxLifetime()

	// create a users table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
			id BIGSERIAL PRIMARY KEY,
			name TEXT
		)`)
	if err != nil {
		panic(err)
	}

	// insert some sample data using a bulk insert
	values := []interface{}{"Alice", "Bob", "Charlie"}
	valueStrings := make([]string, 0, len(values))
	valueArgs := make([]interface{}, 0, len(values))
	for idx, value := range values {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d)", idx+1))
		valueArgs = append(valueArgs, value)
	}

	query := fmt.Sprintf("INSERT INTO users (name) VALUES %s", strings.Join(valueStrings, ", "))
	_, err = db.ExecContext(ctx, query, valueArgs...)
	if err != nil {
		panic(err)
	}

	Users(ctx, db)
	fmt.Println(1)
	Users(ctx, db)
	fmt.Println(2)
	Users(ctx, db)
	fmt.Println(3)
	Users(ctx, db)
	fmt.Println(4)

}

func Users(ctx context.Context, db *sqlx.DB) {
	rows, err := db.QueryContext(ctx, `SELECT * FROM users`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// scan the result into a User struct
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
