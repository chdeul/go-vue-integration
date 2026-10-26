package tools

import (
	"fmt"
	"log"
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

// Current level
var LoggingLevel = InfoLevel

// Stdout  ( blue on vscodium debug output )

func Debug(msg string, args ...any) {
	if LoggingLevel >= DebugLevel { // bleu
		fmt.Printf("Debug: "+msg+"\n", args...)
	}
}
func Info(msg string, args ...any) {
	if LoggingLevel >= InfoLevel { // bleu
		fmt.Printf("Info: "+msg+"\n", args...)
	}
}

// Stderr  ( red on vscodium debug output )

func Log(msg string, args ...any) {
	if LoggingLevel >= LogLevel { // rouge
		log.Printf("Log: "+msg+"\n", args...)
	}
}

func Warn(msg string, args ...any) {
	if LoggingLevel >= WarnLevel { // jaune
		if !Cfg.LogTime {
			log.Printf("\x1b[33mWarn: "+msg+"\x1b[0m\n", args...)
		} else {
			log.Printf("Warn: "+msg+"\n", args...)
		}
	}
}

func Error(msg string, args ...any) {
	if LoggingLevel >= ErrorLevel { // rouge
		log.Printf("Error: "+msg+"\n", args...)
	}
}

func Panic(msg string, args ...any) {
	if LoggingLevel >= PanicLevel { // rouge
		log.Panicf("Panic: "+msg+"\n", args...)
	}
}

func Fatal(msg string, args ...any) {
	if LoggingLevel >= PanicLevel { // rouge
		log.Fatalf("Fatal: "+msg+"\n", args...)
	}
}
