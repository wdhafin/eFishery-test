package db

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// PgDB defines pgdb type
type PgDB struct {
	DB *pg.DB
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	fmt.Println("==================================")
	fmt.Println(q.FormattedQuery())
	return nil
}

// CreatePGConnection return db connection instance
func CreatePGConnection(opts map[string]string) (*PgDB, error) {
	pgdb := pg.Connect(&pg.Options{
		User:     opts["user"],
		Password: opts["password"],
		Database: opts["dbname"],
		Addr:     opts["host"] + ":" + opts["port"],
	})

	var n int
	_, err := pgdb.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		log.Error("Could not connect to PG DB Server:", opts["host"], " at port:", opts["port"])
		log.Fatal(err)
		return nil, err
	}
	log.Info("Connected to PG DB Server: ", opts["host"], " at port:", opts["port"], " successfully!")

	if viper.GetBool(`pg.debug`) {
		pgdb.AddQueryHook(dbLogger{})
	}
	log.Info("DB DEBUG = ", viper.GetBool(`pg.debug`))

	return &PgDB{DB: pgdb}, nil
}
