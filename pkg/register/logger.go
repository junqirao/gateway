package register

import "fmt"

var (
	logTitle      = "[gw_register]"
	defaultLogger = new(logger)
)

// Logger ...
type Logger interface {
	Printf(format string, v ...interface{})
}

type logger struct {
}

func (l logger) Printf(format string, v ...interface{}) {
	fmt.Println(logTitle + " " + fmt.Sprintf(format, v...))
}
