package loggingframework

import "sync"

type Logger struct {
	config *LoggerConfig
	mu     sync.Mutex
}

var instance *Logger
var once sync.Once

func NewLogger() *Logger {
	once.Do(func() {
		instance = &Logger{
			config: &LoggerConfig{
				Level:    Info,
				Appender: &ConsoleAppender{},
			},
		}
	})
	return instance
}

func (l *Logger) SetConfig(config *LoggerConfig) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.config = config
}

func (l *Logger) Log(level LogLevel, message string) error {
	l.mu.Lock()

	if level < l.config.Level {
		defer l.mu.Unlock()
		return nil
	}

	appender := l.config.Appender

	logMessage := &LogMessage{
		Level:   level,
		Message: message,
	}

	defer l.mu.Unlock()
	return appender.Append(logMessage)
}

func (l *Logger) Debug(message string) error {
	return l.Log(Debug, message)
}
func (l *Logger) Info(message string) error {
	return l.Log(Info, message)
}
func (l *Logger) Warn(message string) error {
	return l.Log(Warn, message)
}
func (l *Logger) Error(message string) error {
	return l.Log(Error, message)
}
func (l *Logger) Fatal(message string) error {
	return l.Log(Fatal, message)
}
