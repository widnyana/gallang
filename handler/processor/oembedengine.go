package processor

import (
	"bytes"
	"io/ioutil"

	"github.com/widnyana/oembed"
)

// OEmbedParser will parse the rule and return oembed.Oembed object.
func OEmbedParser() *oembed.Oembed {
	rule, err := ioutil.ReadFile("src/github.com/widnyana/oembed/providers.json")

	if err != nil {
		panic(err)
	}

	svc := oembed.NewOembed()
	svc.ParseProviders(bytes.NewReader(rule))

	return svc
}
