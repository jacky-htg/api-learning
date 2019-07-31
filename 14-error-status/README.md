# Rebel Services

Not every error is an internal server error..

Tasks:
- Add a custom error type that knows about HTTP status codes.
- Make api.Decode return an error with status "400 Bad Request".
- Modify the middleware function to detect this case and use the provided status code.

## File Changes :
Modified services/api/errors.go
Modified services/api/request.go
Modified services/api/response.go
Modified routing/route.go
modified packages/auth/controllers/users.go

## New File :

## Adding Dependency :