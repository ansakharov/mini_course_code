package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

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

	c := make(chan os.Signal)
	signal.Notify(c, os.Kill, os.Interrupt)

	exit := make(chan struct{})
	go func() {
		<-c
		fmt.Printf("signal reached, %v\n", time.Now())

		cancel()

		time.Sleep(time.Second * 10)
		exit <- struct{}{}
	}()

	if err := DoSomeJob(ctx, db); err != nil {
		fmt.Println(err)
	}

	<-exit
	fmt.Println(time.Now())
}

func DoSomeJob(ctx context.Context, db *sqlx.DB) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Printf("request start time: %v\n", time.Now())
			if _, err := db.ExecContext(ctx, "SELECT pg_sleep(5)"); err != nil {
				return err
			}

			fmt.Printf("request end time: %v\n\n", time.Now())
			time.Sleep(time.Second)
		}
	}
}
