package runtime

import (
	"runtime"
)

func GetCurrentFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		return fn.Name()
	}
	return "unknown"
}
