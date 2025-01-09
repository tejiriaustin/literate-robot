package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type ZeroLogger struct {
	logger zerolog.Logger
}

var _ Logger = (*ZeroLogger)(nil)

func NewZeroLogger() *ZeroLogger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &ZeroLogger{logger: logger}
}

func (z *ZeroLogger) convertToZeroFields(fields []Field) []interface{} {
	zeroFields := make([]interface{}, len(fields)*2)
	for i, f := range fields {
		zeroFields[i*2] = f.Key
		zeroFields[i*2+1] = f.Value
	}
	return zeroFields
}

func (z *ZeroLogger) Debug(msg string, fields ...Field) {
	z.logger.Debug().Fields(z.convertToZeroFields(fields)).Msg(msg)
}

func (z *ZeroLogger) Info(msg string, fields ...Field) {
	z.logger.Info().Fields(z.convertToZeroFields(fields)).Msg(msg)
}

func (z *ZeroLogger) Warn(msg string, fields ...Field) {
	z.logger.Warn().Fields(z.convertToZeroFields(fields)).Msg(msg)
}

func (z *ZeroLogger) Error(msg string, fields ...Field) {
	z.logger.Error().Fields(z.convertToZeroFields(fields)).Msg(msg)
}

func (z *ZeroLogger) Fatal(msg string, fields ...Field) {
	z.logger.Fatal().Fields(z.convertToZeroFields(fields)).Msg(msg)
}
