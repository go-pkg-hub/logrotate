# logrotate

go package for creating and rotating log files

## Usage

```go
package main

import (
    "io"
    "github.com/go-pkg-hub/logrotate"
)

func OpenLogFile(file, maxSize string, maxFiles int) (io.WriteCloser, error) {
    opts := []logrotate.Option{
        logrotate.WithMaxSize(logrotate.StringToSize(maxSize)),
        logrotate.WithMaxFiles(maxFiles),
    }
    return logrotate.New(file, opts...)
}

func main() {
    file, err := OpenLogFile("/var/log/app.log", "1m", 3)
    //...
}
```