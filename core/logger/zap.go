package logger

import "go.uber.org/zap"

type ZapLogger struct {
	logger *zap.Logger
}

var _ Logger = (*ZapLogger)(nil)

func NewZapLogger() (*ZapLogger, error) {
	config := zap.NewProductionConfig()
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return &ZapLogger{logger: logger}, nil
}

func (z *ZapLogger) convertToZapFields(fields []Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, f := range fields {
		zapFields[i] = zap.Any(f.Key, f.Value)
	}
	return zapFields
}

func (z *ZapLogger) Debug(msg string, fields ...Field) {
	z.logger.Debug(msg, z.convertToZapFields(fields)...)
}

func (z *ZapLogger) Info(msg string, fields ...Field) {
	z.logger.Info(msg, z.convertToZapFields(fields)...)
}

func (z *ZapLogger) Warn(msg string, fields ...Field) {
	z.logger.Warn(msg, z.convertToZapFields(fields)...)
}

func (z *ZapLogger) Error(msg string, fields ...Field) {
	z.logger.Error(msg, z.convertToZapFields(fields)...)
}

func (z *ZapLogger) Fatal(msg string, fields ...Field) {
	z.logger.Fatal(msg, z.convertToZapFields(fields)...)
}
