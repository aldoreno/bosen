package log

import (
	sglog "github.com/sourcegraph/log"
)

var logger sglog.Logger

func ProvideLogger() sglog.Logger {
	if logger == nil {
		logger = sglog.Scoped("", "root logger")
	}

	return logger
}

/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
