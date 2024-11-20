package postgres

import (
	"fmt"

	"example.com/boiletplate/config"
	"example.com/boiletplate/ent"
)

func NewPsqlDB(config *config.Config) (*ent.Client, error) {

	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.Postgres.Host, config.Postgres.Port, config.Postgres.Username, config.Postgres.Database, config.Postgres.Password, config.Postgres.Sslmode)
	fmt.Printf("config: %s", connString)

	client, err := ent.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	return client, nil
}
