# Go logger Discord
A golang logger send over discord channel
![Preview](https://github.com/fachryansyah/go-logger-discord/blob/master/ss.png)

# Instalation
```bash
$ go get github.com/fachryansyah/go-logger-discord
```

# Usage
Simply setting your auth config for initial logger
```go
package main

import (
    dscrdlog "github.com/fachryansyah/go-logger-discord"
)

func main() {
    log := dscrdlog.Init(dscrdlog.AuthConfig{
        Token: "your token",
        Type:  "Bot",
    }, "your channel id")

    log.Error(
        "a message",
        "a title of error",
        "a description of error",
        "url refrence"
    )
}
```

# API
| Name            | Description               |Return |
| ----------------|---------------------------|-------|
| `log.Error()`   | log an error              | error |
| `log.Info()`    | log an information        | error |
| `log.Success()` | log an something success  | error |
| `log.Fatal()`   | log an fatal error        | error |
| `log.Warn()`    | log an warning message    | error |