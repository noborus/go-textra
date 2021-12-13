# go-textra

[![Go Reference](https://pkg.go.dev/badge/github.com/noborus/go-textra.svg)](https://pkg.go.dev/github.com/noborus/go-textra)

This is a library that translates with [みんなの自動翻訳(minnano-jidou-honyaku)@textra's](https://mt-auto-minhon-mlt.ucri.jgn-x.jp/) API clinet.

You need a [textra](https://mt-auto-minhon-mlt.ucri.jgn-x.jp/) account.

```go
package main

import (
	"fmt"
	"log"

	"github.com/noborus/go-textra"
)

var config = textra.Config{
	ClientID:     "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", // API key
	ClientSecret: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",  // API secret
	Name:         "UserID", // UserID
}

func main() {
	cli,err := textra.New(config)
	if err != nil {
		log.Fatal(err)
	}

	ja, err := cli.Translate(textra.GENERAL_EN_JA, "This is a pen.")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ja) // これはペンです。
}
```
