package log

import (
	"fmt"
	"strings"
	"time"
)

func log(LogLevel string, message string) {
	fmt.Println(fmt.Sprintf("[%s] %s %s", strings.ToUpper(LogLevel), time.Now().Format("2006-01-02 15:04:05"), message))
}
