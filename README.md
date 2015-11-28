# static

[![Build Status](https://secure.travis-ci.org/goforgery/favicon.png?branch=master)](http://travis-ci.org/goforgery/favicon)

Favicon server for Forgery2.

## Use

Caches and serves a `favicon.ico` file if found in the directory where __Forgery2__ is started.

```javascript
package main

import (
	"github.com/goforgery/favicon"
	"github.com/goforgery/forgery2"
)

func main() {
	app := f.CreateApp()
	app.Use(favicon.Create())
	app.Listen(3000)
}
```

## Test

    go test
