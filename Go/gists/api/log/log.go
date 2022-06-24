package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"luvx/config"
	"os"
	"path"
	"time"
)

var stdFormatter *prefixed.TextFormatter  // 命令行输出格式
var fileFormatter *prefixed.TextFormatter // 文件输出格式

func init() {
	stdFormatter = &prefixed.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02.15:04:05.000000",
		ForceFormatting: true,
		ForceColors:     true,
		DisableColors:   false,
	}
	stdFormatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle:  "green",
		WarnLevelStyle:  "yellow",
		ErrorLevelStyle: "red",
		FatalLevelStyle: "41",
		PanicLevelStyle: "41",
		DebugLevelStyle: "blue",
		PrefixStyle:     "cyan",
		TimestampStyle:  "37",
	})

	fileFormatter = &prefixed.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02.15:04:05.000000",
		ForceFormatting: true,
		ForceColors:     false,
		DisableColors:   true,
	}

	logPath := config.Config.GetString("log.logDir")
	logFile := config.Config.GetString("log.mainLog")
	writer, _ := rotatelogs.New(
		path.Join(logPath, logFile),
		rotatelogs.WithLinkName(path.Join(logPath, "main.log")),
		rotatelogs.WithMaxAge(time.Duration(168)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Second),
	)
	writer1, _ := rotatelogs.New(
		path.Join(logPath, config.Config.GetString("log.errorLog")),
		rotatelogs.WithLinkName(path.Join(logPath, "error.log")),
		rotatelogs.WithMaxAge(time.Duration(168)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Second),
	)

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer1,
		log.FatalLevel: writer1,
		log.PanicLevel: writer1,
	}, fileFormatter)

	log.AddHook(lfHook)
	log.SetReportCaller(true)
	log.SetFormatter(stdFormatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func a() {
	log.WithFields(log.Fields{
		"json": "json1",
	}).Info("foo bar")

	log.WithField("json", "json1").Infoln("foo bar")
}
