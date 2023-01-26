package logs

// A log level.
type Level string

// Constants that enumerate log levels.
const (
	// Level for debug or trace information.
	DebugLevel Level = "DEBUG"

	// Level for routine information, such as ongoing status or performance.
	InfoLevel Level = "INFO"

	// Level for normal but significant events, such as start up, shut down, or a
	// configuration change.
	NoticeLevel Level = "NOTICE"

	// Level for warning events that might cause problems.
	WarningLevel Level = "WARNING"

	// Level for error events that are likely to cause problems.
	ErrorLevel Level = "ERROR"

	// Level for critical events that cause more severe problems or outages.
	CriticalLevel Level = "CRITICAL"

	// Level for when a person must take an action immediately.
	AlertLevel Level = "ALERT"

	// Level for when one or more systems are unusable.
	EmergencyLevel Level = "EMERGENCY"
)
