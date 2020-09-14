package init

import (
	"errors"
	"os"

	"github.com/spf13/viper"
	"github.com/wdhafin/eFishery-test/pkg/logger"
)

// setupAuthHelper inits
func setupAuthHelper() {

	if !viper.IsSet("auth.accessSecret") || viper.GetString("auth.accessSecret") == "" {
		logger.CaptureErr(errors.New("auth.accessSecret can not be empty for better security on auth"))
		os.Exit(1)
	}

	if !viper.IsSet("auth.refreshSecret") || viper.GetString("auth.refreshSecret") == "" {
		logger.CaptureErr(errors.New("auth.refreshSecret can not be empty for better security on auth"))
		os.Exit(1)
	}
}
