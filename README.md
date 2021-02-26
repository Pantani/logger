[![Go Reference](https://pkg.go.dev/badge/github.com/Pantani/logger.svg)](https://pkg.go.dev/github.com/Pantani/logger)
[![codecov](https://codecov.io/gh/Pantani/logger/branch/master/graph/badge.svg?token=3CHQF7L76B)](https://codecov.io/gh/Pantani/logger)

# Simple log package

Simple abstraction for logs using [Logrus](https://github.com/sirupsen/logrus).

E.g.:

- Log message:
```go
logger.Info("Loading Observer API")
```

- Log message with params:
```go
logger.Info("Running application", logger.Params{"bind": bind})
```

- Fatal with error:
```go
logger.Fatal("Application failed", err)
```

- The method parameters don't have a sort. You just need to pass them to the method:
```go
logger.Fatal(err, "Application failed")
```

- Create a simple error log:
```go
logger.Error(err)
```

- Create an error log with a message:
```go
logger.Error("Failed to initialize API", err)
```

- Create an error log, with error, message, and params:
```
p := logger.Params{
	"platform": handle,
	"coin":     platform.Coin(),
}
err := platform.Init()
if err != nil {
	logger.Error("Failed to initialize API", err, p)
}
```

- Debug log:
```go
logger.Debug("Loading Observer API")
// OR 
logger.Debug("Loading Observer API", logger.Params{"bind": bind})
```

- Warning log:
```go
logger.Warn("Warning", err)
// OR 
logger.Warn(err, "Warning")
// OR 
logger.Warn("Warning", err, logger.Params{"bind": bind})
```
