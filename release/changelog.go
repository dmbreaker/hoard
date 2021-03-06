package release

import (
	"bytes"
	"text/template"
)

var changeLogTemplate = template.Must(template.New("changelog_template").
	Parse(`# Hoard Changelog{{ range . }}
## Version {{ .Version }}
{{ .Notes }}
{{ end }}`))

func Changelog() string {
	completeChangeLog, err := changelogForReleases(hoardReleases)
	if err != nil {
		panic(err)
	}
	return completeChangeLog
}

func changelogForReleases(rels []Release) (string, error) {
	buf := new(bytes.Buffer)
	err := changeLogTemplate.Execute(buf, rels)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
