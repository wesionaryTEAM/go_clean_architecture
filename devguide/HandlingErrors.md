### `responses` package

The `responses` package provides utility functions to standardize the structure of HTTP responses in the application. It includes functions for handling both success and error responses, ensuring consistency and clarity in API responses.

#### Key Functions:

1. **JSON**: Sends a JSON response with a given status code and data payload.
2. **ErrorJSON**: Sends a JSON response specifically for errors, with a given status code and error message.
3. **SuccessJSON**: Sends a JSON response for successful operations, with a given status code and success message.
4. **JSONWithPagination**: Sends a JSON response with pagination details, including data and pagination metadata like `has_next` and `count`.

#### Error Handling:

The `responses` package also includes functions for handling errors effectively:

1. **HandleValidationError**: Logs and sends a `400 Bad Request` response for validation errors.
2. **HandleErrorWithStatus**: Logs and sends a response with a custom status code for specific errors.
3. **HandleError**: A comprehensive error handler that:
   - Handles custom `APIError` types.
   - Handles `gorm.ErrRecordNotFound` with a `404 Not Found` response.
   - Logs and sends a generic `500 Internal Server Error` response for unhandled errors.
   - Captures unhandled exceptions using Sentry for further analysis.

### Examples for `responses` package

#### Example: Sending a JSON Response
```go
import (
	"github.com/gin-gonic/gin"
	"clean-architecture/pkg/responses"
)

func ExampleHandler(c *gin.Context) {
	data := map[string]string{"message": "Hello, World!"}
	responses.JSON(c, http.StatusOK, data)
}
```

#### Example: Sending an Error Response
```go
import (
	"github.com/gin-gonic/gin"
	"clean-architecture/pkg/responses"
)

func ErrorHandler(c *gin.Context) {
	err := "Something went wrong"
	responses.ErrorJSON(c, http.StatusBadRequest, err)
}
```

#### Example: Sending a Success Response
```go
import (
	"github.com/gin-gonic/gin"
	"clean-architecture/pkg/responses"
)

func SuccessHandler(c *gin.Context) {
	msg := "Operation successful"
	responses.SuccessJSON(c, http.StatusOK, msg)
}
```

### Examples for `HandleValidationError` and `HandleError`

#### Example: Using `HandleValidationError`
```go
import (
	"github.com/gin-gonic/gin"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/responses"
)

func ValidationErrorHandler(c *gin.Context) {
	logger := framework.NewLogger()
	err := errors.New("Invalid input data")
	responses.HandleValidationError(logger, c, err)
}
```

#### Example: Using `HandleError` with `errorz` package
```go
import (
	"github.com/gin-gonic/gin"
	"clean-architecture/pkg/errorz"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/responses"
)

func APIErrorHandler(c *gin.Context) {
	logger := framework.NewLogger()

	// Example of a custom API error
	apiErr := &errorz.APIError{
		StatusCode: 404,
		Message:    "Resource not found",
	}

	responses.HandleError(logger, c, apiErr)
}

func GenericErrorHandler(c *gin.Context) {
	logger := framework.NewLogger()

	// Example of a generic error
	err := errors.New("Something went wrong")
	responses.HandleError(logger, c, err)
}
```

### Defining Common Errors with `errorz` Package

The `errorz` package allows you to define commonly used errors in a centralized manner, making it easier to reuse them across different parts of the application. This approach ensures consistency in error handling and reduces duplication.

#### Steps to Define and Use Common Errors:

1. **Define Common Errors**:
   Use the `errorz` package to define errors that are frequently used, such as validation errors, authentication errors, or resource not found errors.

   ```go
   package errorz

   import "net/http"

   var (
       ErrUnauthorized = &APIError{
           StatusCode: http.StatusUnauthorized,
           Message:    "Unauthorized access",
       }

       ErrResourceNotFound = &APIError{
           StatusCode: http.StatusNotFound,
           Message:    "The requested resource was not found",
       }

       ErrInvalidInput = &APIError{
           StatusCode: http.StatusBadRequest,
           Message:    "Invalid input provided",
       }
   )
   ```

2. **Use Defined Errors in Handlers**:
   Use these predefined errors in your handlers or services to ensure consistent error responses.

   ```go
   import (
       "github.com/gin-gonic/gin"
       "clean-architecture/pkg/errorz"
       "clean-architecture/pkg/framework"
       "clean-architecture/pkg/responses"
   )

   func ExampleHandler(c *gin.Context) {
       logger := framework.NewLogger()

       // Simulate an error condition
       if true { // Replace with actual condition
           responses.HandleError(logger, c, errorz.ErrInvalidInput)
           return
       }

       responses.SuccessJSON(c, http.StatusOK, "Operation successful")
   }
   ```

3. **Benefits of Centralized Error Definitions**:
   - **Consistency**: Ensures that the same error messages and status codes are used across the application.
   - **Maintainability**: Makes it easier to update error messages or status codes in one place.
   - **Reusability**: Reduces duplication by allowing the same error definitions to be reused in multiple places.
