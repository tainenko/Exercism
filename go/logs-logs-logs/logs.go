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
	level := strings.ToLower(s[0])
	level = strings.TrimLeft(level, "[")
	level = strings.TrimRight(level, "]")
	return level
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	message := Message(line)
	level := LogLevel(line)
	return fmt.Sprintf("%s (%s)", message, level)

}
