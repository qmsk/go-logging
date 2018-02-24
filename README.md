# Usage

The zero-valued `logging.Logging` struct functions as a no-op logger that does not output anything.

### `mypackage/logging.go`

```go
package mypackage

import (
    "github.com/qmsk/snmpbot/util/logging"
)

var log logging.Logging

func SetLogging(l logging.Logging) {
    log = l
}
```

### `cmd/mycmd/main.go`

```go
import (
    "github.com/qmsk/snmpbot/util/logging"
    "flag"
    ".../mypackage"
)

var LoggingOptions logging.Options

func init() {
  LoggingOptions.InitFlags()
}

func main() {
  flag.Parse()

  mypackage.SetLogging(LoggingOptions.MakeLogging())

  ...
}
```
