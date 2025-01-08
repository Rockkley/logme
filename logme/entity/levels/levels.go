package levels

type LogLevel int

const (
	Info LogLevel = iota
	Warning
	Debug
	Critical
	All
)

var levelToString = map[LogLevel]string{
	Info:     "INFO",
	Warning:  "WARNING",
	Debug:    "DEBUG",
	Critical: "CRITICAL",
}

func (ll *LogLevel) String() string {
	return levelToString[*ll]
}
