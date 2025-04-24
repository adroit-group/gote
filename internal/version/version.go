package version

import "github.com/adroit-group/go-template/pkg/version"

var Committish string = "unknown"
var BuildDate string = "unknown"

func GetVersion() version.Version {
	return version.Version{
		Committish: Committish,
		BuildDate:  BuildDate,
	}
}
