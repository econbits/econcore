//Copyright (C) 2020  Germán Fuentes Capella

package script

const (
	hDescription = "DESCRIPTION"
	hURL         = "URL"
	hAuthors     = "AUTHORS"
	hLicense     = "LICENSE"
)

var (
	stringHeaders = []string{
		hDescription,
		hURL,
		hLicense,
	}
	listHeaders = []string{
		hAuthors,
	}
)

const (
	defDescription = "Module under construction"
	defUrl         = "https://econbits.org/"
	defLicense     = ""
)

var (
	defAuthors = []string{}
)
