package processor

import (
    "github.com/widnyana/oembed"
    "io/ioutil"
    "bytes"
)

func OEmbedParser() *oembed.Oembed {
    rule, err := ioutil.ReadFile("src/github.com/widnyana/gallang/data/providers.json")

    if err != nil {
        panic(err)
    }

    svc := oembed.NewOembed()
    svc.ParseProviders(bytes.NewReader(rule))

    return svc
}