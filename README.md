# Fiber Structured Access Logging Middleware

This library is created to log Fiber access logs via the structured logging package so it could be easily filtered out based on the "component"
field of the log entry.

In order to use it get the latest version by running:

```shell
go get github.com/iktech/github.com/iktech/fiber-slog-middleware
```

In order to initialise it use the following code snippet:

```go
    app := fiber.New(fiber.Config{
        Prefork:               false,
        AppName:               "My awesome application",
    })

    app.Use(logger.New("access_logger"))
```
