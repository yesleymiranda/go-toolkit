package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if os.Getenv("ENVIRONMENT") == "" {
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano, NoColor: false}

		output.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf("[msg:%s]", i)
		}
		output.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("[%s:", i)
		}
		output.FormatFieldValue = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("%s]", i))
		}

		log.Logger = zerolog.New(output).With().Timestamp().Logger()
	}
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Error(msg string) {
	log.Error().Msg(msg)
}
