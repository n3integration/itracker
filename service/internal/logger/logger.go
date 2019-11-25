/*
 *  Copyright (C) 2019 n3integration
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program. If not, see <https://www.gnu.org/licenses/>.
 */
package logger

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var (
	level string

	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errLogger   *log.Logger
)

func init() {
	flag.StringVar(&level, "level", "debug", "logging level [debug,info,warn,error]")

	debugWriter := ioutil.Discard
	infoWriter := ioutil.Discard
	warnWriter := ioutil.Discard
	errWriter := ioutil.Discard

	switch level {
	case "debug":
		debugWriter = os.Stdout
		infoWriter = os.Stdout
		warnWriter = os.Stdout
		errWriter = os.Stdout
	case "info":
		infoWriter = os.Stdout
		warnWriter = os.Stdout
		errWriter = os.Stdout
	case "warn":
		warnWriter = os.Stdout
		errWriter = os.Stdout
	case "error":
		errWriter = os.Stdout
	}

	debugLogger = log.New(debugWriter, " [debug] ", log.LstdFlags|log.LUTC)
	infoLogger = log.New(infoWriter, " [info ] ", log.LstdFlags|log.LUTC)
	warnLogger = log.New(warnWriter, " [warn ] ", log.LstdFlags|log.LUTC)
	errLogger = log.New(errWriter, " [error] ", log.LstdFlags|log.LUTC)
}

func Debug(message string, args ...interface{}) {
	debugLogger.Printf(message, args...)
}

func Info(message string, args ...interface{}) {
	infoLogger.Printf(message, args...)
}

func Warn(message string, args ...interface{}) {
	warnLogger.Printf(message, args...)
}

func Error(message string, err error) {
	errLogger.Println(message, err)
}

func Fatal(message string, err error) {
	errLogger.Println(message, err)
	os.Exit(1)
}

