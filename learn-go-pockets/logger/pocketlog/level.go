package pocketlog 

// Level represents an available logging level
type Level byte

const (
  // LevelDebug represents the lowest level of log, mostly used to debug
  LevelDebug Level = iota
  // LevelInfo represents the logging level of log, mostly contains information 
  LevelInfo
  // LevelError represents the highest logging level, only used to log errors
  LevelError
)
