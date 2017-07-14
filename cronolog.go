package cronolog

import (
	"log"
	"os"
)

const (
	Ldate         = log.Ldate
	Ltime         = log.Ltime
	Lmicroseconds = log.Lmicroseconds
	Llongfile     = log.Llongfile
	Lshortfile    = log.Lshortfile
	LUTC          = log.LUTC
	LstdFlags     = log.LstdFlags
)

const (
	TimeMinuteFormat = "200601021504"
	TimeHourFormat   = "2006010215"
	TimeDayFormat    = "20060102"
	DefaultFormat    = TimeDayFormat
)

type logger struct {
	l  *log.Logger
	fd *os.File
}

func (l *logger) SetRotate() {

}

func (l *logger) SetRotateFromat(format string) {

}