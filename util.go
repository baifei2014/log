package log

import (
	"runtime"
	"strconv"
	"sync"
)

var fm sync.Map

// funcName get func name.
func funcName(skip int) (name string) {
	if pc, _, lineNo, ok := runtime.Caller(skip); ok {
		if v, ok := fm.Load(pc); ok {
			name = v.(string)
		} else {
			name = runtime.FuncForPC(pc).Name() + ":" + strconv.FormatInt(int64(lineNo), 10)
			fm.Store(pc, name)
		}
	}
	return
}
