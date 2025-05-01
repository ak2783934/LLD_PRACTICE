package loggingframework

type ConsoleAppender struct {
}

func NewConsoleAppender() *ConsoleAppender {
	return &ConsoleAppender{}
}

func (ca *ConsoleAppender) Append(message *LogMessage) error {
	// Print the log message to the console
	println(message.String())
	return nil
}
