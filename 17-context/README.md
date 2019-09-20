# Rebel Services

Context

Long running operations should be given a deadline. The idiomatic way to handle cancellation is passing context.Context to functions that know to check for cancellation and terminate early.

Task
- Add context.Context argument to all function in packages/auth/models/user.go
- Pass the ctx variable to db.QueryContext, db.QueryRowContext, db.PrepareContext and stmt.ExecContext in packages/auth/models/user.go
- Pass the value of r.Context() from packages/auth/controllers/users.go every call service/function at packages/auth/models/user.go
- Pass the value of context.Background() from packages/auth/models/models_test/user_test.go every call service/function at packages/auth/models/user.go 

## File Changes :
- packages/auth/models/user.go
- packages/auth/controllers/users.go
- packages/auth/models/models_test/user_test.go

## New File :

## Adding Dependency :