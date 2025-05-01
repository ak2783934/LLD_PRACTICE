package loggingframework

import "fmt"

type LogMessage struct {
	Level     LogLevel
	Message   string
	TimeStamp string
}

func (lm *LogMessage) String() string {
	return fmt.Sprintf("%s [%s] %s", lm.TimeStamp, lm.Level.String(), lm.Message)
}
