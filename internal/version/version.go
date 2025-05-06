package version

import "github.com/adroit-group/gote/pkg/version"

var (
	committish = "unknown"
	buildDate  = "unknown"
)

func GetVersion() version.Version {
	return version.Version{
		Committish: committish,
		BuildDate:  buildDate,
	}
}
