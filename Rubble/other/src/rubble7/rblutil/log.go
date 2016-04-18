/*
Copyright 2013-2016 by Milo Christiansen

This software is provided 'as-is', without any express or implied warranty. In
no event will the authors be held liable for any damages arising from the use of
this software.

Permission is granted to anyone to use this software for any purpose, including
commercial applications, and to alter it and redistribute it freely, subject to
the following restrictions:

1. The origin of this software must not be misrepresented; you must not claim
that you wrote the original software. If you use this software in a product, an
acknowledgment in the product documentation would be appreciated but is not
required.

2. Altered source versions must be plainly marked as such, and must not be
misrepresented as being the original software.

3. This notice may not be removed or altered from any source distribution.
*/

package rblutil

import "math/rand"
import "time"
import "fmt"
import "os"
import "io"
import "bytes"

// Logger writes to rubble.log and os.Stdout by default, you can add more locations if you wish.
type Logger struct {
	io.Writer
	
	WC int
	WB *bytes.Buffer
	
	// This buffer stores the log up to now, useful for the web UI mostly.
	LB *bytes.Buffer
}

// NewLogger creates a new logger.
func NewLogger(other ...io.Writer) (error, *Logger) {
	// We never close this file (probably not a good idea, but where would we close it?)
	file, err := os.Create("./rubble.log")
	if err != nil {
		return err, nil
	}

	lb := new(bytes.Buffer)
	
	writers := make([]io.Writer, 3, len(other)+3)
	writers[0] = os.Stdout
	writers[1] = file
	writers[2] = lb
	writers = append(writers, other...)
	return nil, &Logger{io.MultiWriter(writers...), 0, new(bytes.Buffer), lb}
}

// Printf prints to the log, see ftm.Printf.
func (log *Logger) Printf(format string, msg ...interface{}) {
	fmt.Fprintf(log, format, msg...)
}

// Println prints to the log, see ftm.Println.
func (log *Logger) Println(msg ...interface{}) {
	fmt.Fprintln(log, msg...)
}

// Print prints to the log, see ftm.Print.
func (log *Logger) Print(msg ...interface{}) {
	fmt.Fprint(log, msg...)
}

// Warnf prints to the log and the warnings buffer, see ftm.Printf.
func (log *Logger) Warnf(format string, msg ...interface{}) {
	fmt.Fprintf(log, format, msg...)
	fmt.Fprintf(log.WB, format, msg...)
	log.WC++
}

// Warnln prints to the log and the warnings buffer, see ftm.Println.
func (log *Logger) Warnln(msg ...interface{}) {
	fmt.Fprintln(log, msg...)
	fmt.Fprintln(log.WB, msg...)
	log.WC++
}

// Warn prints to the log and the warnings buffer, see ftm.Print.
func (log *Logger) Warn(msg ...interface{}) {
	fmt.Fprint(log, msg...)
	fmt.Fprint(log.WB, msg...)
	log.WC++
}

// WarnExtraf prints to the log and the warnings buffer but does not increment the warning count, see ftm.Printf.
func (log *Logger) WarnExtraf(format string, msg ...interface{}) {
	fmt.Fprintf(log, format, msg...)
	fmt.Fprintf(log.WB, format, msg...)
}

// WarnExtraln prints to the log and the warnings buffer but does not increment the warning count, see ftm.Println.
func (log *Logger) WarnExtraln(msg ...interface{}) {
	fmt.Fprintln(log, msg...)
	fmt.Fprintln(log.WB, msg...)
}

// WarnExtra prints to the log and the warnings buffer but does not increment the warning count, see ftm.Print.
func (log *Logger) WarnExtra(msg ...interface{}) {
	fmt.Fprint(log, msg...)
	fmt.Fprint(log.WB, msg...)
}

// WarnOnlyf prints to the warnings buffer, see ftm.Printf.
func (log *Logger) WarnOnlyf(format string, msg ...interface{}) {
	fmt.Fprintf(log.WB, format, msg...)
	log.WC++
}

// WarnOnlyln prints to the warnings buffer, see ftm.Println.
func (log *Logger) WarnOnlyln(msg ...interface{}) {
	fmt.Fprintln(log.WB, msg...)
	log.WC++
}

// WarnOnly prints to the warnings buffer, see ftm.Print.
func (log *Logger) WarnOnly(msg ...interface{}) {
	fmt.Fprint(log.WB, msg...)
	log.WC++
}

// WarnExtraf prints to the warnings buffer but does not increment the warning count, see ftm.Printf.
func (log *Logger) WarnOnlyExtraf(format string, msg ...interface{}) {
	fmt.Fprintf(log.WB, format, msg...)
}

// WarnExtraln prints to the warnings buffer but does not increment the warning count, see ftm.Println.
func (log *Logger) WarnOnlyExtraln(msg ...interface{}) {
	fmt.Fprintln(log.WB, msg...)
}

// WarnExtra prints to the warnings buffer but does not increment the warning count, see ftm.Print.
func (log *Logger) WarnOnlyExtra(msg ...interface{}) {
	fmt.Fprint(log.WB, msg...)
}

// ClearWarnings clears the warnings buffer and resets the warning count.
func (log *Logger) ClearWarnings() {
	log.WB.Reset()
	log.WC = 0
}

// Separator writes a section separator to the log.
// Use for consistency.
func (log *Logger) Separator() {
	fmt.Fprint(log, "==============================================\n")
}

// Header writes the standard header to the log.
// Use for consistency.
func (log *Logger) Header(version string) {
	rand.Seed(time.Now().Unix())
	fmt.Fprint(log, "Rubble v"+version+"\n")
	fmt.Fprint(log, startupLines[rand.Int()%(len(startupLines)-1)]+"\n")
	log.Separator()
}

var startupLines = [...]string{
	"After Blast comes Rubble.",
	"Modding made easy!",
	"Scriptable!",
	"Templates!",
	"Now with random startup lines!",
	"Why did I add this feature?",
	"Rubblize it!",
	"Now with a web UI!",
	"Use the command line!",
	"Configurable!",
	"Now with more addons you will never use!",
	"Please report any problems.",
	"Feedback is greatly appreciated!",
	"Unintentionally Ironic!",
	"Why do these all end with exclamation points!",
	"Free exclamation points!",
	"Guaranteed 50% bug free!",
	"There better not be an error log!",
	"Run it again, this line might change.",
	"Over 100 addons!",
	"Now with meta data!",
	"Lua Scripting!",
	"Read the documentation!",
	"Runs natively on Windows, Linux, and OSX!",
	"Under continuous development since June 2013!",
	"Open source!",
	"Supports DFHack!",
	"50+ KB log files!",
	"Include the ENTIRE log in any bug reports!",
	"Long hours of debugging later...",
	"Rule-based tileset installation!",
}
