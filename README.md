# go-textra

This is a library that translates with [みんなの自動翻訳(minnano-jidou-honyaku)@textra's](https://mt-auto-minhon-mlt.ucri.jgn-x.jp/) API.

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
	cli := textra.New(config)

	ja, err := cli.Translate(textra.GENERAL_EN_JA, "This is a pen.")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ja) // これはペンです。
}
```
