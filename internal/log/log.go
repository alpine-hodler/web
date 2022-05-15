package log

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/alpine-hodler/sdk/internal/env"
)

type Level uint8

const (
	INFO Level = iota
	WARN
	DEBUG
)

var (
	currentLevel  *Level
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger

	// suppress will only be set true of there is some init() error that isn't worth terminating the runtmie over but also
	// prevents us from using the logging functionality.  An example of this is when there is no $HOME environment
	// variable, in which case we suppress logging.
	suppress bool

	homeNotDefined = fmt.Errorf("$HOME is not defined")
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		if err != homeNotDefined {
			// if home is not defined, then do not log.
			// TODO: Check to see if this behavior is mac-specific.
			suppress = true
			return
		}
		log.Fatal(err)
	}

	// TODO: (1) make the logPath OS-specific, right not it's specific to MacOS
	// TODO: (2) make the path an optional env variable
	// TODO: (3) see if we can use a bazel directory to log data when running a build.  Bazel won't know what the $HOME
	// TODO:     environment variable is
	logDir := path.Join(home, "Library/Logs/AlpineHodler")
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, fs.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	// TODO: Make the name of the log an optional environment variable.
	logPath := path.Join(logDir, "sdk.log")
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, fs.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	debugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func getLogLevel() Level {
	if currentLevel != nil {
		return *currentLevel
	}

	envLevel := env.AlpineHodlerLogLevel.Get()
	if envLevel == "" {
		info := INFO
		currentLevel = &info
		return *currentLevel
	}

	raw, err := strconv.Atoi(env.AlpineHodlerLogLevel.Get())
	if err != nil {
		log.Fatal(err)
	}
	rawLevel := Level(raw)
	currentLevel = &rawLevel
	return *currentLevel
}

// Debug logs debug data to the log file.
func Debug(msg string, args ...interface{}) {
	if suppress {
		return
	}
	if getLogLevel() >= DEBUG {
		debugLogger.Println(fmt.Sprintf(msg, args...))
	}
}

// Info logs informational data to the log file.
func Info(msg string, args ...interface{}) {
	if suppress {
		return
	}
	infoLogger.Println(fmt.Sprintf(msg, args...))
}

// Warn logs warning data to the log file.
func Warn(msg string, args ...interface{}) {
	if suppress {
		return
	}
	warningLogger.Println(fmt.Sprintf(msg, args...))
}

// Logf will takes a level and data to log to that level.
func Logf(level Level, msg string, args ...interface{}) {
	if suppress {
		return
	}
	switch level {
	case DEBUG:
		Debug(msg, args...)
	case INFO:
		Info(msg, args...)
	case WARN:
		Warn(msg, args...)
	}
}
