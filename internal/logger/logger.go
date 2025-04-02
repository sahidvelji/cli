package logger

import (
	"path/filepath"

	"github.com/pterm/pterm"
)

// Logger provides methods for logging different types of messages
type Logger interface {
	// Println logs a message without logging level
	Println(message string)
	// Info logs general information
	Info(message string)
	// Success logs successful operations
	Success(message string)
	// Warning logs warnings
	Warning(message string)
	// Error logs errors
	Error(message string)
	// Debug logs debug information (only when debug mode is enabled)
	Debug(message string)
	// SetDebug enables or disables debug mode
	SetDebug(enabled bool)
	// IsDebugEnabled returns whether debug mode is enabled
	IsDebugEnabled() bool
	// FileCreated logs a file creation event
	FileCreated(path string)
	// FileFailed logs a file creation failure
	FileFailed(path string, err error)
	// GenerationStarted logs the start of a generation process
	GenerationStarted(generatorType string)
	// GenerationComplete logs the completion of a generation process
	GenerationComplete(generatorType string)
}

// DefaultLogger is the default implementation of Logger
type DefaultLogger struct {
	debugEnabled bool
}

// New creates a new DefaultLogger
func New() *DefaultLogger {
	return &DefaultLogger{
		debugEnabled: false,
	}
}

// SetDebug enables or disables debug mode
func (l *DefaultLogger) SetDebug(enabled bool) {
	l.debugEnabled = enabled
	if enabled {
		pterm.EnableDebugMessages()
	}
}

// IsDebugEnabled returns whether debug mode is enabled
func (l *DefaultLogger) IsDebugEnabled() bool {
	return l.debugEnabled
}

// Println logs a message without logging level
func (l *DefaultLogger) Println(message string) {
	pterm.Println(message)
}

// Info logs general information
func (l *DefaultLogger) Info(message string) {
	pterm.Info.Println(message)
}

// Success logs successful operations
func (l *DefaultLogger) Success(message string) {
	pterm.Success.Println(message)
}

// Warning logs warnings
func (l *DefaultLogger) Warning(message string) {
	pterm.Warning.Println(message)
}

// Error logs errors
func (l *DefaultLogger) Error(message string) {
	pterm.Error.Println(message)
}

// Debug logs debug information (only when debug mode is enabled)
func (l *DefaultLogger) Debug(message string) {
	if l.debugEnabled {
		pterm.Debug.Println(message)
	}
}

// FileCreated logs a file creation event
func (l *DefaultLogger) FileCreated(path string) {
	prettyPath := pterm.LightWhite(filepath.Clean(path))
	pterm.Success.Printf("Created %s\n", prettyPath)
}

// FileFailed logs a file creation failure
func (l *DefaultLogger) FileFailed(path string, err error) {
	prettyPath := pterm.LightWhite(filepath.Clean(path))
	pterm.Error.Printf("Failed to create %s: %v\n", prettyPath, err)
}

// GenerationStarted logs the start of a generation process
func (l *DefaultLogger) GenerationStarted(generatorType string) {
	pterm.Info.Printf("Generating a typesafe client for %s\n", generatorType)
}

// GenerationComplete logs the completion of a generation process
func (l *DefaultLogger) GenerationComplete(generatorType string) {
	pterm.Success.Printf("Successfully generated client. Happy coding!\n")
}

// Default is a singleton instance of DefaultLogger
var Default Logger = New()
