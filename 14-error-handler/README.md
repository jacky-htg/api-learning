# Rebel Services

Error Handler

Not every error is an internal server error. We will handle all error with standard format response { status_code, status_message, data }

Task
- Add a custom error type that knows about HTTP status codes
- Make api.Decode return an error with status "400 Bad Request"
- Change response using standard format [ ResponseFormat ]
- All call api.Response varying with new response format 

## File Changes :
- Modified packages/auth/controllers/users.go
- Modified libraries/api/request.go
- Modified libraries/api/response.go

## New File :
- libraries/api/error.go
- libraries/api/status_code.go
- libraries/api/status_message.go

## Adding Dependency :