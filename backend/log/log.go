package log

import (
	"bosen/manifest"
	sglog "github.com/sourcegraph/log"
)

func InitLogger() *sglog.PostInitCallbacks {
	return sglog.Init(sglog.Resource{
		Name:    manifest.AppName,
		Version: manifest.CommitHash,
	})
}

/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
