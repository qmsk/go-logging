# Usage

The zero-valued `logging.Logging` struct functions as a no-op logger that does not output anything.

### `mypackage/logging.go`

```go
package client

import (
    "github.com/qmsk/snmpbot/util/logging"
)

var log logging.Logging

func SetLogging(l logging.Logging) {
    log = l
}
```
