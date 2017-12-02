[![Build Status](https://travis-ci.org/alexbyk/panicif.svg?branch=master)](https://travis-ci.org/alexbyk/panicif)
# About

Instead of writing
```go
err := IsPositive(-1)
if err != nil {
  panic(err)
}
```

use a shorthand:
```go
panicif.Err(IsPositive(-1))
```


# Installation
```
go get -u github.com/alexbyk/panicif
```
# Usage

```go
package main

import (
  "fmt"

  "github.com/alexbyk/panicif"
)

func IsPositive(val int) error {
  if val < 0 {
    return fmt.Errorf("%d isn't a positive integer", val)
  }
  return nil
}

// will panic
func main() {
  panicif.Err(IsPositive(-1))
}

```

# Copyright
Copyright 2018, [alexbyk.com](https://alexbyk.com)
