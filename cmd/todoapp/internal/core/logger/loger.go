package core_logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger

	file *os.File
}

func NewLogger(config Config) (*Logger, error) {
	level := zap.NewAtomicLevel()

	if err := level.UnmarshalText([]byte(config.Level)); err != nil {
		return nil, fmt.Errorf("unmarshal log level: %w", err)
	}

	if err := os.MkdirAll(config.Folder, 0755); err != nil {
		return nil, fmt.Errorf("create log directory: %w", err)
	}

	timestamp := time.Now().
		UTC().
		Format("2006-01-02T15-04-05")

	logPath := filepath.Join(
		config.Folder,
		fmt.Sprintf("log-%s.log", timestamp),
	)

	file, err := os.OpenFile(
		logPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)

	if err != nil {
		return nil, fmt.Errorf("open log file: %w", err)
	}

	zapConfig := zap.NewDevelopmentEncoderConfig()
	zapConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04::05:00000")
	
	zapEncoder := zapcore.NewConsoleEncoder(zapConfig)

	core := zapcore.NewTee(
	zapcore.NewCore(
		zapEncoder,
		zapcore.AddSync(os.Stdout),
		level,
	),
	zapcore.NewCore(
		zapEncoder,
		zapcore.AddSync(file),
		level,
	),
)

zaplogger := zap.New(core, zap.AddCaller())

return &Logger{
	Logger: zaplogger,
	file:   file,
}, nil
	}

func (l *Logger) Close() error {
	if err := l.Sync(); err != nil {
		return err
	}

	return l.file.Close()
}