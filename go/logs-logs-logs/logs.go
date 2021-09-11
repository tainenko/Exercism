package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Message extracts the message from the provided log line.
func Message(line string) string {
	s := strings.Split(line, ":")
	return strings.TrimSpace(s[1])
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	return utf8.RuneCountInString(Message(line))
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	s := strings.Split(line, ":")
	logLevel := s[0]
	return strings.ToLower(logLevel[1 : len(logLevel)-1])
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	message := Message(line)
	logLevel := LogLevel(line)
	return fmt.Sprintf("%s (%s)", message, logLevel)

}
