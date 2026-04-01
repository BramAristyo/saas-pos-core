package logger

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type GormZapLogger struct {
	ZapLogger     *zap.Logger
	SlowThreshold time.Duration
}

func NewGormZapLogger(zapLogger *zap.Logger, slowThreshold time.Duration) logger.Interface {
	return &GormZapLogger{
		ZapLogger:     zapLogger,
		SlowThreshold: slowThreshold,
	}
}

func (l *GormZapLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *GormZapLogger) Info(ctx context.Context, msg string, data ...any) {
	l.ZapLogger.Sugar().Infof(msg, data...)
}

func (l *GormZapLogger) Warn(ctx context.Context, msg string, data ...any) {
	l.ZapLogger.Sugar().Warnf(msg, data...)
}

func (l *GormZapLogger) Error(ctx context.Context, msg string, data ...any) {
	l.ZapLogger.Sugar().Errorf(msg, data...)
}

func (l *GormZapLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	if err != nil && err != logger.ErrRecordNotFound {
		l.ZapLogger.Error("Database Error",
			zap.Error(err),
			zap.Duration("elapsed", elapsed),
			zap.Int64("rows", rows),
			zap.String("sql", sql),
		)
		return
	}

	if elapsed > l.SlowThreshold && l.SlowThreshold != 0 {
		l.ZapLogger.Warn("Slow Query Detected",
			zap.Duration("elapsed", elapsed),
			zap.Int64("rows", rows),
			zap.String("sql", sql),
		)
	}
}
