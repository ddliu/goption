# goption

[![Build Status](https://travis-ci.org/ddliu/goption.png?branch=master)](https://travis-ci.org/ddliu/goption)

Option container for golang

##  Installation

```bash
go get github.com/ddliu/goption
```

## Usage

```go
import "github.com/ddliu/goption"

// create option

o := goption.NewOption(map[string]interface{} {
    "timeout": 100,
    "cacheable": true,
})

o.MergeMap(map[string]interface{} {
    "prefix": "key_",
    "size": 1.5,
})

timeout, ok := o.GetInt("timeout")

cachable := o.MustGetBool("cacheable")

size: = o.MustGetFloat32("size")

xyz, ok := o.GetInt("xyz") // ok is false

xyz := o.MustGetInt("xyz") // panic!
```

## Changelog

### v0.1.0 (2013-09-23)

First release

## License

Licensed under the MIT license

