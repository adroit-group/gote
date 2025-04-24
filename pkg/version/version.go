package version

// Version represents the version information of the application.
type Version struct {
	// Commitish is the tag or commit hash of the given application version.
	//
	// See: https://git-scm.com/docs/gitglossary#Documentation/gitglossary.txt-aiddefcommit-ishacommit-ishalsocommittish
	Committish string `json:"committish"`
	// BuildDate is the date when the application was built.
	BuildDate string `json:"build_date"`
}

// VersionProvider is a function that returns the version information of the application.
type VersionProvider func() Version
