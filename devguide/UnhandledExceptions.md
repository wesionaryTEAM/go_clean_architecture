### Handling Unhandled Exceptions with Sentry

Unhandled exceptions can occur in any application and need to be captured and logged effectively to ensure they are addressed promptly. In this project, we use Sentry to capture unhandled exceptions and provide detailed context for debugging.

#### Capturing Unhandled Exceptions

The `utils.SendSentryMsg` function can be used to send custom error messages to Sentry. For unhandled exceptions, you can use the `utils.CurrentSentryService.CaptureException` method to capture the exception and send it to Sentry.

Example:

```go
import (
	"errors"
	"github.com/gin-gonic/gin"
	"clean-architecture/pkg/utils"
)

func ExampleUnhandledExceptionHandler(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			err := errors.New("Unhandled exception occurred")
			utils.CurrentSentryService.CaptureException(err)
			c.JSON(500, gin.H{"error": "An unexpected error occurred. Please try again later."})
		}
	}()

	// Simulate an unhandled exception
	panic("Simulated panic")
}
```


#### Best Practices

1. Use `defer` and `recover` to handle panics and capture unhandled exceptions.
2. Pass relevant context to Sentry to make debugging easier.
3. Ensure sensitive information is not sent to Sentry.