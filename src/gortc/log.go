// Copyright (C) AlexWoo(Wu Jie) wj19840501@gmail.com
//
// RTC Log

package main

import (
	"os"
	"rtclib"
	"strconv"

	"github.com/alexwoo/golib"
)

type logctx struct {
	logLevel int
	log      *golib.Log
}

func (ctx *logctx) Prefix() string {
	return "[main] " + strconv.Itoa(os.Getpid())
}

func (ctx *logctx) Suffix() string {
	return ""
}

func (ctx *logctx) LogLevel() int {
	return ctx.logLevel
}

var mainlogCtx *logctx

func initLog() {
	logPath := rtclib.FullPath(config.LogFile)
	logLevel := golib.LoglvEnum.ConfEnum(config.LogLevel, golib.LOGINFO)

	mainlogCtx = &logctx{
		logLevel: logLevel,
		log:      golib.NewLog(logPath),
	}

	if mainlogCtx.log == nil {
		os.Exit(1)
	}

	return
}

func LogDebug(format string, v ...interface{}) {
	mainlogCtx.log.LogDebug(mainlogCtx, format, v...)
}

func LogInfo(format string, v ...interface{}) {
	mainlogCtx.log.LogInfo(mainlogCtx, format, v...)
}

func LogError(format string, v ...interface{}) {
	mainlogCtx.log.LogError(mainlogCtx, format, v...)
}

func LogFatal(format string, v ...interface{}) {
	mainlogCtx.log.LogFatal(mainlogCtx, format, v...)
}
