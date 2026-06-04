# Handling Errors

Error handling is split across two packages:

- **`errorz`** — defines the canonical application error type (`APIError`) and a set of
  reusable, predefined errors.
- **`responses`** — serializes errors (and successes) into a consistent HTTP response shape
  and provides handler helpers that translate Go errors into `APIError`s.

## The `APIError` type (`errorz`)

```go
type APIError struct {
    Code       string         // machine-readable identifier (UPPER_SNAKE); the client localizes from this
    StatusCode int            // HTTP status for transport; never serialized into the body
    Severity   Severity       // info | warn | error
    Message    string         // optional human-readable message set by the backend; omitted when empty
    Params     map[string]any // optional context / i18n interpolation values; omitted when empty
    Cause      error          // internal wrapped error for logging/Sentry/Unwrap; never serialized
}
```

`Severity` is one of:

```go
errorz.SeverityInfo  // "info"
errorz.SeverityWarn  // "warn"
errorz.SeverityError // "error"
```

### Creating errors

Use `New`. The message is optional — omit it to let the client localize from `Code`:

```go
// no default message
errorz.New(errorz.CodeInvalidUUID, http.StatusBadRequest, errorz.SeverityWarn)

// with an optional default message
errorz.New(errorz.CodeInvalidUUID, http.StatusBadRequest, errorz.SeverityWarn, "Invalid UUID")
```

### Builder methods (clone-on-write)

The builders never mutate the receiver, so it is safe to derive from the shared
package-level errors. Each returns a new `*APIError`:

```go
errorz.ErrBadRequest.WithMessage("Email is already taken")     // set a custom message
errorz.ErrBadRequest.WithParam("field", "email")               // set a single param
errorz.ErrBadRequest.WithParams(map[string]any{"field": "email"}) // merge params
errorz.ErrInternal.Wrap(err)                                   // attach an internal cause (not serialized)
```

They chain:

```go
return errorz.ErrConflict.
    WithMessage("User already exists").
    WithParam("email", user.Email)
```

### Predefined errors

Reusable errors are declared centrally so codes, statuses, and severities stay consistent:

- `pkg/errorz/base.go` — canonical HTTP errors (`ErrBadRequest`, `ErrNotFound`,
  `ErrInternal`, …) and derived ones (`ErrAlreadyExists`, `ErrSomethingWentWrong`). Their
  default message is the standard `http.StatusText(...)`.
- `pkg/errorz/common_errors.go` — cross-cutting semantic errors (`ErrRecordNotFound`,
  `ErrInvalidToken`, `ErrUnauthorizedAccess`, …).
- Domain packages declare their own (e.g. `domain/user/api_error.go` →
  `ErrInvalidUserID`, `ErrUserAlreadyExists`).

### `From`

`From` converts an arbitrary error into an `*APIError`, returning it unchanged if it already
is (or wraps) one, and falling back to `ErrInternal` otherwise:

```go
api := errorz.From(err)
```

## Response shape (`responses`)

### Success

```go
responses.Success(c, http.StatusOK, data, nil)            // {"success": true, "data": ...}
responses.PaginationSuccess(c, http.StatusOK, data, total) // adds {"meta": {"pagination": {...}}}
```

### Error envelope

`responses.Error` writes the `APIError` using its `StatusCode` as the HTTP status. The body is:

```json
{
  "error": {
    "code": "USER_ALREADY_EXISTS",
    "severity": "warn",
    "params": { "email": "a@b.com" },
    "message": "User already exists"
  }
}
```

`params` and `message` are omitted when empty. `StatusCode` and `Cause` are never included in
the body.

## Handler helpers (`responses`)

| Function | Behavior |
|---|---|
| `HandleError(logger, c, err)` | If `err` is/wraps an `*APIError`, writes it as-is. Else if it is/wraps `gorm.ErrRecordNotFound`, returns `404`. Otherwise logs the error, captures it to Sentry, and returns a generic `500`. |
| `HandleErrorWithParams(logger, c, apiErr, params)` | Writes `apiErr` with the given `params` attached. |
| `HandleErrorWithStatus(logger, c, statusCode, err)` | Logs `err` and returns it under `CodeCustomError` with the given status (severity `error` and a Sentry capture when `statusCode >= 500`). |
| `HandleValidationError(logger, c, err)` | Passes an `*APIError` through unchanged (no error log). For ozzo `validation.Errors`, returns `400` with a `validation_errors` param (sorted by field). Otherwise delegates to `HandleError`. |

### Example: central error handling

```go
func (u *Controller) GetUserByID(c *gin.Context) {
    uid, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        responses.HandleValidationError(u.logger, c, ErrInvalidUserID) // 400, code INVALID_USER_ID
        return
    }

    user, err := u.service.GetUserByID(uint(uid))
    if err != nil {
        responses.HandleError(u.logger, c, err) // gorm not-found → 404; anything else → 500 + Sentry
        return
    }

    responses.Success(c, http.StatusOK, user, nil)
}
```

### Example: returning a known error with context and a custom message

```go
if exists {
    responses.Error(c, ErrUserAlreadyExists.
        WithParam("email", user.Email).
        WithMessage("User already exists"))
    return
}
```

## Why centralized error definitions

- **Consistency** — the same code, status, and severity are reused everywhere.
- **Maintainability** — change a code/status/severity in one place.
- **Localization-friendly** — clients translate from `code` (+ `params`); `message` is an
  optional backend override, not a requirement.
