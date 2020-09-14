package init

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// setupLogrus inits the way logrus logs our app
func setupLogrus() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

}
