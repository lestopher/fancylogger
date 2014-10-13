package fancylogger

import (
	"io"
	"log"

	"github.com/mgutz/ansi"
)

var (
	// Success is the color of success
	Success = ansi.ColorCode("green+bh")
	// Info is the color of information
	Info = ansi.ColorCode("white+bh")
	// Error is the color of a meltdown
	Error = ansi.ColorCode("red+bh")
	// Warning means gun it
	Warning = ansi.ColorCode("yellow+bh")
	// Reset puts the colors back the way it was
	Reset = ansi.ColorCode("reset")
)

// FancyLogger will give you nicely printed color logs
type FancyLogger struct {
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

// NewFancyLogger will set up a fancy logger instance, but requires all of
// outputs to be defined
func NewFancyLogger(traceHandle io.Writer, infoHandle io.Writer,
	warningHandle io.Writer, errorHandle io.Writer) *FancyLogger {
	return &FancyLogger{
		Trace: log.New(traceHandle, "TRACE: ",
			log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
		Info: log.New(infoHandle, Info+"INFO: "+Reset,
			log.Ldate|log.Ltime|log.Lmicroseconds),
		Warning: log.New(warningHandle, Warning+"WARNING: "+Reset,
			log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
		Error: log.New(errorHandle, Error+"ERROR: "+Reset,
			log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
	}
}
