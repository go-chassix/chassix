package log

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"

	"pgxs.io/panguxs/pkg/chassis/config"
)

// var formatter = &zt_formatter.ZtFormatter{
// 	CallerPrettyfier: func(f *runtime.Frame) (string, string) {
// 		filename := path.Base(f.File)
// 		return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
// 	},
// 	Formatter: nested.Formatter{
// 		// HideKeys:    true,
// 		FieldsOrder:     []string{"component", "category"},
// 		TimestampFormat: "2006-01-02 15:04:05",
// 	},
// }

// log := logrus.New()
var formatter = &nested.Formatter{
	HideKeys:        true,
	FieldsOrder:     []string{"component", "category"},
	TimestampFormat: "2006-01-02 15:04:05",
}

func init() {

}

//Logger custom logger
type Logger struct {
	*log.Logger
}

//Entry log entry
type Entry struct {
	// *Logger
	*log.Entry
}

//New new logger
func New() *Logger {
	nLog := log.New()
	lg := &Logger{
		Logger: nLog,
	}
	level := log.Level(config.Logging().Level)
	if level >= log.DebugLevel {
		formatter.HideKeys = false
	}
	nLog.SetFormatter(formatter)
	nLog.SetLevel(level)
	nLog.SetReportCaller(true)
	// nLog.SetReportCaller(config.Logging().ReportCaller)

	return lg
}

//Service setting svc name
func (l *Logger) Service(svc string) *Entry {
	return &Entry{l.WithField("component", svc)}
}

//Category setting svc name
func (l *Logger) Category(name string) *Entry {
	return &Entry{l.WithField("category", name)}
}

//SetReqInfo setting svc name
func (e *Entry) SetReqInfo(req *restful.Request) *Entry {
	return &Entry{e.WithFields(log.Fields{
		"req":   req.Request.Method + " " + req.Request.URL.String(),
		"reqId": req.Attribute("reqId"),
	})}
}

//Category setting svc name
func (e *Entry) Category(name string) *Entry {
	return &Entry{e.WithField("category", name)}
}

var defaultLogger = &Logger{
	log.StandardLogger(),
}

//StdLogger std logger
func StdLogger() *Logger {
	return defaultLogger
}
