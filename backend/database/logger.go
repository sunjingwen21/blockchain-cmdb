package database

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm/logger"
)

// CustomLogger is a custom GORM logger
type CustomLogger struct{}

func (l *CustomLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *CustomLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	log.Printf("[GORM] INFO: "+msg, data...)
}

func (l *CustomLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	log.Printf("[GORM] WARN: "+msg, data...)
}

func (l *CustomLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	log.Printf("[GORM] ERROR: "+msg, data...)
}

func (l *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		log.Printf("[GORM] ERROR: %s [%.3fms] rows:%d err:%v", sql, float64(elapsed.Milliseconds()), rows, err)
	} else {
		log.Printf("[GORM] QUERY: %s [%.3fms] rows:%d", sql, float64(elapsed.Milliseconds()), rows)
	}
}
