package logx

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/emicklei/go-restful/v3"
	log "github.com/sirupsen/logrus"

	"c6x.io/chassix.v2/config"
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
	if config.NotNil() {
		level := log.Level(config.Logging().Level)
		if level >= log.DebugLevel {
			formatter.HideKeys = false
		}
		nLog.SetLevel(level)
		nLog.SetReportCaller(config.Logging().ReportCaller)
		formatter.NoColors = config.Logging().NoColors
		if config.Logging().ReportCaller {
			formatter.CallerFirst = config.Logging().CallerFirst
		}
	}

	nLog.SetFormatter(formatter)
	return lg
}

const (
	fieldSvcKey = "service"
	fieldComKey = "component"
	fieldCatKey = "category"
)

//Service setting svc name
func (l *Logger) Service(svc string) *Entry {
	return &Entry{l.WithField(fieldSvcKey, svc)}
}

//Component setting svc name
func (l *Logger) Component(com string) *Entry {
	return &Entry{l.WithField(fieldComKey, com)}
}

//Category setting svc name
func (l *Logger) Category(name string) *Entry {
	return &Entry{l.WithField(fieldCatKey, name)}
}

//SetReqInfo setting svc name
func (e *Entry) SetReqInfo(req *restful.Request) *Entry {
	return &Entry{e.WithFields(log.Fields{
		"req":   req.Request.Method + " " + req.Request.URL.String(),
		"reqId": req.Attribute("reqId"),
	})}
}

//Component setting svc name
func (e *Entry) Component(com string) *Entry {
	return &Entry{e.WithField(fieldComKey, com)}
}

//Category setting svc name
func (e *Entry) Category(name string) *Entry {
	return &Entry{e.WithField(fieldCatKey, name)}
}

var defaultLogger = &Logger{
	log.StandardLogger(),
}

//StdLogger std logger
func StdLogger() *Logger {
	return defaultLogger
}
