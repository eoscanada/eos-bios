package bios

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logger struct {
	OutputFile   io.Writer
	OutputScreen io.Writer
	Debug        bool
}

func NewLogger() *Logger {
	fl, err := os.Create("output.log")
	if err != nil {
		log.Fatalln("Couldn't open output.log:", err)
	}

	return &Logger{
		OutputFile:   fl,
		OutputScreen: os.Stdout,
	}
}

func (l *Logger) Debugln(args ...interface{}) {
	if l.Debug {
		fmt.Fprintln(l.OutputScreen, args...)
	}
	fmt.Fprintln(l.OutputFile, args...)
}

func (l *Logger) Println(args ...interface{}) {
	fmt.Fprintln(l.OutputScreen, args...)
	fmt.Fprintln(l.OutputFile, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.Debug {
		fmt.Fprintf(l.OutputScreen, format, args...)
	}
	fmt.Fprintf(l.OutputFile, format, args...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	fmt.Fprintf(l.OutputScreen, format, args...)
	fmt.Fprintf(l.OutputFile, format, args...)
}
