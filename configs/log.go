package configs

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strings"
)

type logFormatter struct {
	logrus.TextFormatter
}

func (f logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	prettyCaller := func(frame *runtime.Frame) string {
		_, fileName := filepath.Split(frame.File)
		return fmt.Sprintf("%s:%d", fileName, frame.Line)
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	b.WriteString(fmt.Sprintf("[%s] %s", entry.Time.Format(f.TimestampFormat), strings.ToUpper(entry.Level.String())))
	if entry.HasCaller() {
		b.WriteString(fmt.Sprintf("[%s]", prettyCaller(entry.Caller)))
	}
	b.WriteString(fmt.Sprintf("[%s]", entry.Message))
	return b.Bytes(), nil
}
