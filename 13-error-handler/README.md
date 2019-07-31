# Rebel Services

Handling application errors in a consistent and reliable way is very repetitive. Provide some support for that in the web package.

Tasks:
- Define a custom signature for all handler functions that includes returning errors.
- Add a middleware function that will be ran for all handlers which deals with the returned errors.


## File Changes :
- Modified packages/auth/controllers/users.go
- Modified routing/route.go

## New File :
- services/api/errors.go

## Adding Dependency :