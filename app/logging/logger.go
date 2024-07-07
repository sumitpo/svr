// logger.go

package logging

import (
	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	// Log formatter
	Log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
		CallerFirst: true,
	})
	// Log level
	Log.Level = logrus.InfoLevel
	// Output to stdout by default
	Log.Out = os.Stdout
}
