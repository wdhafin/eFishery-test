package init

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wdhafin/eFishery-test/internal/db"
	"github.com/wdhafin/eFishery-test/pkg/utils"
)

// ConnectToPGServer is a function to init PostgreSQL connection
func ConnectToPGServer() (*db.PgDB, error) {

	if utils.IsProductionEnv() && (!viper.IsSet("pg.password") || viper.GetString("pg.password") == "") {
		logrus.Error("pg.password can not be empty!")
		os.Exit(1)
	}

	dbpg, err := db.CreatePGConnection(map[string]string{
		"host":     viper.GetString(`pg.host`),
		"port":     viper.GetString(`pg.port`),
		"user":     viper.GetString(`pg.user`),
		"password": viper.GetString(`pg.password`),
		"dbname":   viper.GetString(`pg.dbname`),
	})
	if err != nil {
		os.Exit(1)
	}
	return dbpg, err
}
