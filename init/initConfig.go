package init

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/wdhafin/eFishery-test/pkg/utils"
)

// setupMainConfig loads app config to viper
func setupMainConfig() {
	logrus.Info("Executing init/config")

	viper.SetConfigFile("config/app.yml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Info("err: ", err)
	}

	viper.SetConfigFile("config/main.yml")
	err = viper.MergeInConfig()
	if err != nil {
		logrus.Info("err: ", err)
	}

	if utils.IsFileExist(".env.yml") {
		logrus.Info("Local .env.yml file is found, now assigning it with default config")
		viper.SetConfigFile(".env.yml")
		err = viper.MergeInConfig()
		if err != nil {
			logrus.Info("err: ", err)
		}
	}

	viper.SetEnvPrefix(`app`)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

	logrus.Info("Config- APP_ENV: ", utils.GetEnv())

}
