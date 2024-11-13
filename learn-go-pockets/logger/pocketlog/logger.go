package pocketlog

// Logger is used to log information
type Logger struct{

}

// Debugf formats and prints a message if the log level is Debug
func (l *Logger) Debugf(format string, args ...any) {
  // code goes here
}

func (l *Logger) Infof(format string, args ...any) {
  // code goes here
}

func (l *Logger) Errorf(format string, args ...any) {
  // code goes here
}
