package manifest

var AppVersion string
var BuildTime string
var CommitMsg string
var CommitHash string
var ReleaseVersion string

type Manifest struct {
	AppVersion     string
	BuildTime      string
	CommitMsg      string
	CommitHash     string
	ReleaseVersion string
}

func Info() Manifest {
	return Manifest{
		AppVersion,
		BuildTime,
		CommitMsg,
		CommitHash,
		ReleaseVersion,
	}
}
