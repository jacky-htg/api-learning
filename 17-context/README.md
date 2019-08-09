# Rebel Services

Context

Long running operations should be given a deadline. The idiomatic way to handle cancellation is passing context.Context to functions that know to check for cancellation and terminate early.

Tasks:
- Add context.Context argument to all function in models/user.go 
- Pass the ctx variable to db.SelectContext, db.GetContext, db.PreparexContext and stmt.ExecContext in models/user.go 
- Pass the value of r.Context() from controllers/users.go every call service/function at models/user.go
- Pass the value of context.Background() from models/models_test/user_test.go every call service/function at models/user.go

## File Changes :
- models/user.go
- controllers/users.go
- models/models_test/user_test.go

## New File :

## Adding Dependency :