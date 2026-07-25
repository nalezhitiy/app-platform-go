package logger

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type GormLogger struct {
	log           *zap.Logger
	level         gormlogger.LogLevel
	slowThreshold time.Duration
}

func NewGormLogger(
	log *zap.Logger,
	debug bool,
) gormlogger.Interface {

	level := gormlogger.Warn

	if debug {
		level = gormlogger.Info
	}

	return &GormLogger{
		log:           log,
		level:         level,
		slowThreshold: 200 * time.Millisecond,
	}
}

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	cp := *l
	cp.level = level
	return &cp
}

// Не логуємо внутрішні Info-повідомлення GORM.
// Весь SQL приходить через Trace().
func (l *GormLogger) Info(
	context.Context,
	string,
	...any,
) {
}

func (l *GormLogger) Warn(
	ctx context.Context,
	msg string,
	data ...any,
) {
	if l.level < gormlogger.Warn {
		return
	}

	l.log.Warn(
		fmt.Sprintf(msg, data...),
	)
}

func (l *GormLogger) Error(
	ctx context.Context,
	msg string,
	data ...any,
) {
	if l.level < gormlogger.Error {
		return
	}

	l.log.Error(
		fmt.Sprintf(msg, data...),
	)
}

func (l *GormLogger) Trace(
	ctx context.Context,
	begin time.Time,
	fc func() (string, int64),
	err error,
) {

	if l.level == gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)

	sql, rows := fc()

	fields := []zap.Field{
		zap.Duration("elapsed", elapsed),
		zap.Int64("rows", rows),
		zap.String("SQL", sql),
	}

	// ---------- ERROR ----------
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {

		if l.level >= gormlogger.Error {

			fields = append(fields, zap.Error(err))

			l.log.Error(
				"SQL",
				fields...,
			)
		}

		return
	}

	// ---------- SLOW QUERY ----------
	if elapsed >= l.slowThreshold {

		if l.level >= gormlogger.Warn {

			l.log.Warn(
				"SLOW SQL",
				fields...,
			)
		}

		return
	}

	// ---------- NORMAL QUERY ----------
	if l.level >= gormlogger.Info {

		l.log.Info(
			"SQL",
			fields...,
		)
	}
}
