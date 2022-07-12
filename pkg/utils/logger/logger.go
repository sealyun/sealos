// Copyright © 2022 sealos.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	defaultLogger *zap.Logger
)

// init default logger with only console output info above
func init() {
	zc := zapcore.NewTee(newConsoleCore(zap.InfoLevel))
	defaultLogger = zap.New(zc)
}

// CfgConsoleLogger config for console logs
// cfg donot support concurrent calls (as any package should init cfg at startup once)
func CfgConsoleLogger(debugMode bool, showPath bool) {
	level, zos := genConfigs(debugMode, showPath)

	zc := zapcore.NewTee(newConsoleCore(level))

	defaultLogger = zap.New(zc, zos...)
}

//TODO: export more file configs
// CfgConsoleAndFileLogger config for both console and file logs
// cfg donot support concurrent calls (as any package should init cfg at startup once)
func CfgConsoleAndFileLogger(debugMode bool, logDir, name string, showPath bool) {
	level, zos := genConfigs(debugMode, showPath)

	filename := fmt.Sprintf("%s/%s.log", logDir, name)

	zc := zapcore.NewTee(newConsoleCore(level), newFileCore(filename, level))

	defaultLogger = zap.New(zc, zos...)
}

func genConfigs(debugMode bool, showPath bool) (zapcore.LevelEnabler, []zap.Option) {
	level := zapcore.InfoLevel
	if debugMode {
		level = zapcore.DebugLevel
	}

	zos := []zap.Option{
		// zap.AddStacktrace(zapcore.WarnLevel),
	}
	if showPath {
		// skip self wrapper
		zos = append(zos, zap.AddCaller(), zap.AddCallerSkip(2))
	}

	return level, zos
}

func newConsoleCore(le zapcore.LevelEnabler) zapcore.Core {
	consoleLogger := zapcore.Lock(os.Stdout)

	zec := zap.NewProductionEncoderConfig()
	zec.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	zec.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(zec)

	return zapcore.NewCore(consoleEncoder, consoleLogger, le)
}

func newFileCore(filename string, le zapcore.LevelEnabler) zapcore.Core {
	//TODO: export more rotate configs
	fileLogger := zapcore.AddSync(&lumberjack.Logger{
		Filename: filename,
		MaxSize:  10, // megabytes per file
	})

	zec := zap.NewProductionEncoderConfig()
	zec.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(zec)
	return zapcore.NewCore(fileEncoder, fileLogger, le)
}

// IsDebugMode check DebugLevel enabled
func IsDebugMode() bool {
	return defaultLogger.Core().Enabled(zapcore.DebugLevel)
}

// Fatal logs a message at emergency level and exit.
func Fatal(f interface{}, v ...interface{}) {
	defaultLogger.Sugar().Fatalf(formatLog(f, v...))
}

// Panic logs a message at emergency level and exit.
func Panic(f interface{}, v ...interface{}) {
	defaultLogger.Sugar().Panicf(formatLog(f, v...))
}

// Error logs a message at error level.
func Error(f interface{}, v ...interface{}) {
	defaultLogger.Sugar().Errorf(formatLog(f, v...))
}

// Warn logs a message at warning level.
func Warn(f interface{}, v ...interface{}) {
	defaultLogger.Sugar().Warnf(formatLog(f, v...))
}

// Info logs a message at info level.
func Info(f interface{}, v ...interface{}) {
	defaultLogger.Sugar().Infof(formatLog(f, v...))
}

// Debug logs a message at debug level.
func Debug(f interface{}, v ...interface{}) {
	defaultLogger.Sugar().Debugf(formatLog(f, v...))
}

func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f := f.(type) {
	case string:
		msg = f
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
