// Package grad : grad provides feature for feed detection from html document
package grad

import (
	"io"

	"gopkg.in/xmlpath.v2"
)

// DetectFeed from html
func DetectFeed(body io.Reader) (feeds []string) {
	html, err := xmlpath.ParseHTML(body)
	if err != nil {
		return
	}

	var paths []*xmlpath.Path
	paths = append(paths, xmlpath.MustCompile(`//link[@rel='alternate' and @type='application/rss+xml']/@href`))
	paths = append(paths, xmlpath.MustCompile(`//link[@rel='alternate' and @type='application/atom+xml']/@href`))

	for _, path := range paths {
		iter := path.Iter(html)
		for iter.Next() {
			feeds = append(feeds, iter.Node().String())
		}
	}

	return
}
