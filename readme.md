# Go logger Discord
A golang logger send over discord channel

# Instalation
```bash
$ go get github.com/fachryansyah/go-logger-discord
```

# Usage
Simply setting your auth config for initial logger
```go
log := Init(AuthConfig{
	Token: "your token",
	Type:  "Bot",
}, "your channel id")
```
call logger inside your middleware or on error trace
```go
log.Error(
    "a message",
    "a title of error",
    "a description of error",
    "url refrence"
)
```
# API
| Name            | Description               |Return |
| ----------------|---------------------------|-------|
| `log.Error()`   | log an error              | error |
| `log.Info()`    | log an information        | error |
| `log.Success()` | log an something success  | error |
| `log.Fatal()`   | log an fatal error        | error |
| `log.Warn()`    | log an warning message    | error |