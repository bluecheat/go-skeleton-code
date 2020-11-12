package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
)

const (
	LFhostname   = "hostname"
	LFstatus     = "status"
	LFmethod     = "method"
	LFip         = "ip"
	LFremote     = "remote"
	LFlatency    = "latency"
	LFdataLength = "dataLength"
	LFuserAgent  = "userAgent"
	LFurlPath    = "urlPath"
	LFlogFunc    = "LogFunc"
	LFRequestID  = "RequestID"
)

//CodeLineNumberHook 로그가 찍히는 코드라인 위치 출력을 위한 Hook
type CodeLineNumberHook struct{}

//Levels CodeLineNumberHook이 적용되는 로그레벨: 현재 전부, 추후 Error나 Debug시에만 출력하려면 변경 필요
func (h *CodeLineNumberHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

//Fire CodeLineNumberHook 구현 메소드
func (h *CodeLineNumberHook) Fire(entry *logrus.Entry) error {
	// WithFields 여부에 관계 없이 정확한 위치가 출력되도록 처리
	// https://github.com/sirupsen/logrus/issues/63 참고

	pc := make([]uintptr, 4, 4)
	cnt := runtime.Callers(7, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		funcName := fu.Name()
		if !strings.Contains(funcName, "skeleton-code/logger") &&
			!strings.Contains(funcName, "github.com/sirupsen/logrus") &&
			!strings.Contains(funcName, "skeleton-code/server.Logger") {
			file, line := fu.FileLine(pc[i] - 1)
			entry.Data[LFlogFunc] = fmt.Sprintf("%s:%d:%s", path.Base(file), line, path.Ext(funcName)[1:])
			break
		}
	}
	return nil
}

// 형태

type LogFormatter struct {
	hostname string
	color    bool
	once     sync.Once
}

func (f *LogFormatter) init() error {
	hostname, _ := os.Hostname()
	f.hostname = hostname
	return nil
}

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return terminal.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}

//Format 로그 포맷팅 수행
func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	f.once.Do(func() {
		f.color = checkIfTerminal(entry.Logger.Out) && (runtime.GOOS != "windows")
		f.init()
	})

	isColored := f.isColored()
	b.WriteString(entry.Time.Format("2006-01-02 15:04:05.999"))
	b.WriteString(" [")

	level := strings.ToUpper(entry.Level.String())
	if isColored {
		f.writeColorString(entry, b, level[0:4])
	} else {
		f.appendValue(b, level[0:4], level[0:4])
	}
	b.WriteString("] ")
	b.WriteString(f.hostname)

	b.WriteByte(' ')
	f.appendValue(b, entry.Data[LFlogFunc], "")
	b.WriteString("\t")
	// server log
	if entry.Data[LFmethod] != nil {
		b.WriteString("[HTTP] ")
		f.appendValue(b, entry.Data[LFmethod], "")
		b.WriteByte(' ')
		f.appendValue(b, entry.Data[LFurlPath], "")
		b.WriteByte(' ')
		f.appendValue(b, entry.Data[LFstatus], "")
		b.WriteByte(' ')
		f.appendValue(b, entry.Data[LFlatency], "")
		b.WriteByte(' ')
		f.appendValue(b, entry.Data[LFuserAgent], "")
	} else {
		b.WriteString("[APP]")
	}
	b.WriteString(" | ")

	if isColored {
		f.writeColorString(entry, b, entry.Message)
	} else {
		f.appendValue(b, entry.Message, level)
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *LogFormatter) appendValue(b *bytes.Buffer, value interface{}, nilVal string) {
	if value == nil {
		b.WriteString(nilVal)
	} else {
		stringVal, ok := value.(string)
		if !ok {
			stringVal = fmt.Sprint(value)
		}

		b.WriteString(stringVal)
	}
}

func (f *LogFormatter) isColored() bool {
	return f.color
}

func (f *LogFormatter) writeColorString(entry *logrus.Entry, b *bytes.Buffer, value interface{}) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = 37 // gray
	case logrus.WarnLevel:
		levelColor = 33 // yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}
	if value == nil {
		value = ""
	}
	b.Write([]byte(fmt.Sprintf("\x1b[%dm%s\x1b[0m", levelColor, value)))
}
