package tools

import (
	"fmt"
	"log"
	"os"
)

// Simple log manager with levels

const (
	NoneLevel int = iota
	PanicLevel
	ErrorLevel
	WarnLevel
	InfoLevel // Default
	LogLevel
	DebugLevel
)

var Rouge, Jaune, Magenta, Reset string

// Current level
var LoggingLevel = InfoLevel

func init() {
	log.SetOutput(os.Stdout)
}

func SetLogLevel(level string) {

	// Positionne le niveau de log
	switch level {
	case "None":
		LoggingLevel = NoneLevel
	case "Panic":
		LoggingLevel = PanicLevel
	case "Error":
		LoggingLevel = ErrorLevel
	case "Warn":
		LoggingLevel = WarnLevel
	case "Info":
		LoggingLevel = InfoLevel
	case "Log":
		LoggingLevel = LogLevel
	case "Debug":
		LoggingLevel = DebugLevel
	case "Default":
	case "":
		level = "Default"
	default:
		Warn("Niveau de log inconnu : %s", level)
		level = "Default"
	}
	if level != "Default" {
		Info("Niveau de log : %s", level)
	}
}

// Colorize logs
func LogWithColor(active bool) {
	if active {
		Rouge = "\x1b[31m"
		Jaune = "\x1b[33m"
		Magenta = "\x1b[35m"
		Reset = "\x1b[0m"
	} else {
		Rouge = ""
		Jaune = ""
		Magenta = ""
		Reset = ""
	}
}

// Stdout  ( blue on vscodium debug output )

func Debug(msg string, args ...any) {
	if LoggingLevel >= DebugLevel { // bleu
		log.Printf("Debug: "+msg+"\n", args...)
	}
}
func Info(msg string, args ...any) {
	if LoggingLevel >= InfoLevel { // bleu
		log.Printf("Info : "+msg+"\n", args...)
	}
}

// Stderr  ( red on vscodium debug output )

func Log(msg string, args ...any) {
	if LoggingLevel >= LogLevel { // rouge
		log.Printf("Log  : "+msg+"\n", args...)
	}
}

func Warn(msg string, args ...any) {
	if LoggingLevel >= WarnLevel { // jaune
		log.Printf(Jaune+"Warn : "+msg+Reset+"\n", args...)
	}
}

func Error(msg string, args ...any) {
	if LoggingLevel >= ErrorLevel { // rouge
		log.Printf(Rouge+"Error: "+msg+Reset+"\n", args...)
	}
}

func Panic(msg string, args ...any) {
	if LoggingLevel >= PanicLevel { // rouge
		log.Panicf(Magenta+"Panic: "+msg+Reset+"\n", args...)
	}
}

func Fatal(msg string, args ...any) {
	if LoggingLevel >= PanicLevel { // rouge
		log.Fatalf(Magenta+"Fatal: "+msg+Reset+"\n", args...)
	}
}

// Checks a value and log error message
func CheckError(val error, args ...any) bool {
	if val != nil {
		return Check(val, args...)
	}
	return true
}

// Checks a value and log error message
func Check(val any, args ...any) bool {

	switch v := val.(type) {
	case error:
		if v != nil {
			showError(val, args...)
			return false
		}
	case string:
		if v == "" {
			showError(val, args...)
			return false
		}
	case bool:
		if !v {
			showError(val, args...)
			return false
		}
	case int:
		if v < 0 {
			showError(val, args...)
			return false
		}
	default: // empty array or nil object
		if v == nil {
			showError(val, args...)
			return false
		}
	}
	return true
}

// Checks a value and log error message
func AssertError(val error, args ...any) bool {
	if val != nil {
		return CheckError(val, args...)
	}
	return true
}

// Check and panic
func Assert(val any, args ...any) bool {

	if !Check(val, args...) {
		var ok bool
		var err error
		if err, ok = val.(error); !ok {
			err = fmt.Errorf("Assert failed %s", fmt.Sprint(val))
		}
		panic(err)
	}
	return true
}

// Catch Assert or other panic  ( defer tools.Recover() )
func Recover() {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			Warn("Recover error %v", err)
		}
	}
}

// Log un message d'erreur de check  formaté
func showError(val any, args ...any) {

	var ok bool
	var msg string
	if len(args) > 1 {
		if msg, ok = args[0].(string); ok {
			msg = fmt.Sprintf(msg, args[1:]...)
		} else {
			msg = "Check failed "
			for _, a := range args {
				msg += " " + fmt.Sprint(a)
			}
			args = nil
		}
	}
	if msg != "" {
		Error(msg)
	} else {
		Error("Check failed %s", fmt.Sprint(val))
	}
}
