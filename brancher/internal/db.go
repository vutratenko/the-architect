package internal

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	log "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func getConn() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), viper.GetString("dbstring"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func postBranch(branch string) {
	conn := getConn()
	defer conn.Close(context.Background())
	var branchName string
	err := conn.QueryRow(context.Background(), "INSERT INTO branches (name, lifetime_hours) VALUES ($1, 24) RETURNING name;", branch).Scan(&branchName)
		if err != nil {
			log.Error().Err(err)
		}
	if err != nil {
			log.Error().Err(err)
		}
}