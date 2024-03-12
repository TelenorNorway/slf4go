# Simple Logging Facade fo Go (SLF4Go)

This module aims to provide a simple logging facade for Go. It is inspired by the
[Simple Logging Facade for Java (SLF4J)](http://www.slf4j.org/). The goal is to
provide a simple and efficient logging facade that can be used in Go applications
to provide a consistent logging experience.

## Usage

Get module.

```shell
go get github.com/telenornorway/slf4go
```

On application start, before calling any of the functions you must initialize a driver to use. This is done by calling
`slf4go.UseDriver` with a driver that implements the `slf4go.Driver` interface. You can use a drive like
[Telelog](https://github.com/telenornorway/telelog) to get started.

<!-- @formatter:off -->
```go
package main

import "github.com/telenornorway/slf4go"

func main() {
    var myDriver slf4go.Driver = some_driver.GetDriverSomehow()
    slf4go.UseDriver(myDriver)
    
    var log = slf4go.GetLogger()
    
    log.Info("Hello, %s!", "World")
}
```
<!-- @formatter:on -->

Setting up Telelog would look something like this:

<!-- @formatter:off -->
```go
package main

import (
    "github.com/telenornorway/slf4go"
    "github.com/telenornorway/telelog"
)

func main() {
    telelog.Initialize(slf4go.UseDriver, telelog.DefaultConfig)
    
    var log = slf4go.GetLogger()
    
    log.Info("Hello, %s!", "World")
}
```
