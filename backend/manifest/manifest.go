package manifest

// Following fields are automatically fields on build time "-ldflags".
var AppName string
var AppVersion string
var BuildTime string
var CommitMsg string
var CommitHash string
var ReleaseVersion string

type Manifest struct {
	AppName        string
	AppVersion     string
	BuildTime      string
	CommitMsg      string
	CommitHash     string
	ReleaseVersion string
}

func Info() Manifest {
	return Manifest{
		AppName,
		AppVersion,
		BuildTime,
		CommitMsg,
		CommitHash,
		ReleaseVersion,
	}
}
