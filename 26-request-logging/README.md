# Rebel Services

Request Logging.

Tasks:
- Create middleware to log each request on middleware/logger.go
- Manipulate request with new values and passing it down through context
- Ensure api.Response updates the value.

## File Changes :
- libraries/api/app.go
- middleware/error.go
- libraries/api/response.go
- routing/route.go
- controllers/access.go
- controllers/auths.go
- controllers/checks.go
- controllers/roles.go
- controllers/users.go

## New File :
- middleware/logger.go

## Adding Dependency :