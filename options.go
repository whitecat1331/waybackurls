package waybackurls

type Options struct {
	domains         []string
	dates           bool
	noSubs          bool
	getVersionsFlag bool
}

func CreateOptions(domains []string, dates bool,
	noSubs bool, getVersionsFlag bool) *Options {
	return &Options{
		domains:         domains,
		dates:           dates,
		noSubs:          noSubs,
		getVersionsFlag: getVersionsFlag,
	}
}

func CreateDefaultOptions(domains []string) *Options {
	return CreateOptions(domains, false, false, false)
}
