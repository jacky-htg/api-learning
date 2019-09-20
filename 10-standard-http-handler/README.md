# Rebel Services

Standard Http Handler

httprouter is not using standard http handle with 3 params, http.ResponseWriter,  *http.Request, and httprouter.Params. In this chapter, we will change httprouter.Handle using standard http.Handle

Task:
- Change httprouter.Handle with http.Handle in libraries/api/app.go
- Change each controller with remove httprouter.Params

## File Changes :
- Modified libraries/api/app.go
- Modified packages/auth/controllers/users.go

## New File :

## Adding Dependency :