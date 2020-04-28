package log

import (
	"fmt"
	golog "log"
	"os"
	"path/filepath"
	"time"

	dconf "auth-demo-backend/lib/conf"
)

// Config log config
type Config struct {
	FilePath string `json:"file_path"`
	// RotateMinutes int64  `json:"rotate_minutes"`
	// RotateSuffix  string `json:"ratate_suffix"`
}

var LogFlag = golog.LstdFlags | golog.Llongfile

// NewLog new log instance
func NewLog(conf *Config) (dlog *Log, err error) {
	if !filepath.IsAbs(conf.FilePath) {
		conf.FilePath = filepath.Join(dconf.Root, conf.FilePath)
	}

	dlog = &Log{
		infoLog: &golog.Logger{},
		warnLog: &golog.Logger{},
		rotator: time.NewTicker(time.Second),

		infoFile: conf.FilePath,
		warnFile: conf.FilePath + ".wf",
	}

	dlog.infoLog.SetPrefix("[INFO] ")
	dlog.warnLog.SetPrefix("[WARN] ")

	dlog.infoWriter, dlog.warnWriter, err = dlog.genWriter()
	if err != nil {
		return
	}

	dlog.infoLog.SetOutput(dlog.infoWriter)
	dlog.infoLog.SetFlags(LogFlag)
	dlog.warnLog.SetOutput(dlog.warnWriter)
	dlog.warnLog.SetFlags(LogFlag)
	go dlog.rotate()

	return dlog, nil
}

// Log log object
type Log struct {
	infoLog *golog.Logger
	warnLog *golog.Logger

	infoWriter *os.File
	warnWriter *os.File

	infoFile string
	warnFile string

	rotator *time.Ticker
	// rotateSuffix string

	config *Config
}

func (l *Log) String() string {
	return fmt.Sprintf("info_file: %s, warn_file: %s", l.infoFile, l.warnFile)
}

var perm os.FileMode = 0644
var logFlag int = os.O_WRONLY | os.O_APPEND | os.O_CREATE

func (l *Log) genWriter() (infoWriter *os.File, warnWriter *os.File, err error) {
	if !fileExist(filepath.Dir(l.infoFile)) {
		if err = os.MkdirAll(filepath.Dir(l.infoFile), 0777); err != nil {
			return
		}
	}

	infoWriter, err = os.OpenFile(l.infoFile, logFlag, perm)
	if err != nil {
		return
	}

	warnWriter, err = os.OpenFile(l.warnFile, logFlag, perm)
	return
}

func (l *Log) rotate() {
	for {
		select {
		case <-l.rotator.C:
			if time.Now().Minute() == 0 {
				if err := l.doRotate(); err != nil {
					golog.Printf("rotate fail: %v \n", err)
				}
			}
		}
	}
}

func (l *Log) doRotate() (err error) {
	now := time.Now()
	rotateName := fmt.Sprintf("%s.%s", l.infoFile, now.Format("2006010215"))

	// Dont consider file existed
	if err = os.Rename(l.infoFile, rotateName); err != nil {
		return
	}
	rotateName = fmt.Sprintf("%s.%s", l.warnFile, now.Format("2006010215"))

	// Dont consider file existed
	err = os.Rename(l.warnFile, rotateName)

	i, w, e := l.genWriter()
	if e != nil {
		return e
	}

	l.infoLog.SetOutput(i)
	l.infoLog.SetFlags(LogFlag)
	l.infoWriter.Close()
	l.infoWriter = i

	l.warnLog.SetOutput(w)
	l.warnLog.SetFlags(LogFlag)
	l.warnWriter.Close()
	l.warnWriter = w

	return
}

func fileExist(path string) bool {
	if _, err := os.Stat(path); os.IsExist(err) {
		return true
	}
	return false
}

// Info info log
func (l *Log) Info(v ...interface{}) {
	l.infoLog.Output(2, fmt.Sprint(v...))
}

// Warn warn log
func (l *Log) Warn(v ...interface{}) {
	l.warnLog.Output(2, fmt.Sprint(v...))

}

// Info info log
func Info(v ...interface{}) {
	demo.infoLog.Output(3, fmt.Sprintln(v...))
}

// Warn warn log
func Warn(v ...interface{}) {
	demo.warnLog.Output(3, fmt.Sprintln(v...))
}

var demo *Log

// Init init package
func Init(config *Config) (err error) {
	demo, err = NewLog(config)
	if demo != nil {
		golog.Printf("log init: %v", demo)
	}

	return
}
