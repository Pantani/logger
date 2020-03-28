# Simple log package

Simple abstraction for logs using [Logrus](https://github.com/sirupsen/logrus).

E.g.:

- Log message:
`logger.Info("Loading Observer API")`

- Log message with params:
`logger.Info("Running application", logger.Params{"bind": bind})`

- Fatal with error:
`logger.Fatal("Application failed", err)`

- The method parameters don't have a sort. You just need to pass them to the method:
`logger.Fatal(err, "Application failed")`

- Create a simple error log:
`logger.Error(err)`

- Create an error log with a message:
`logger.Error("Failed to initialize API", err)`

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
`logger.Debug("Loading Observer API")`
 or 
`logger.Debug("Loading Observer API", logger.Params{"bind": bind})`

- Warning log:
`logger.Warn("Warning", err)`
 or 
`logger.Warn(err, "Warning")`
 or 
`logger.Warn("Warning", err, logger.Params{"bind": bind})`
