In Go CLI applications, logging is typically done using the standard library's `log` package or by using more advanced logging libraries like `logrus` or `zap`. Here's how you can use each approach:

### Using the Standard Library's `log` Package:

```go
package main

import (
	"log"
	"os"
)

func main() {
	// Open a log file for writing
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	defer logFile.Close()

	// Set the log output to the log file
	log.SetOutput(logFile)

	// Log some messages
	log.Println("This is a log message")
	log.Printf("This is a formatted log message with a variable: %s", "value")
	log.Fatalf("This is a fatal log message")
}
```

In this example:
- We open a log file `app.log` for writing.
- Set the output of the `log` package to this log file.
- Use functions like `log.Println`, `log.Printf`, or `log.Fatalf` to log messages to the file.

### Using a Third-Party Logging Library (e.g., `logrus`):

First, you need to install the library:

```bash
go get github.com/sirupsen/logrus
```

Then, you can use it in your code:

```go
package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Create a new instance of logrus logger
	logger := log.New()

	// Set the log level
	logger.SetLevel(log.DebugLevel)

	// Open a log file for writing
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	defer logFile.Close()

	// Set the log output to the log file
	logger.SetOutput(logFile)

	// Log some messages
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
}
```

In this example:
- We create a new instance of `logrus.Logger`.
- Set the log level to control which messages are logged.
- Open a log file and set the log output to this file.
- Use methods like `logger.Info`, `logger.Warn`, `logger.Error` to log messages at different levels.

Choose the approach that best fits your needs and the complexity of your logging requirements. The standard library's `log` package is sufficient for simple logging needs, while third-party libraries like `logrus` offer more features and flexibility.
