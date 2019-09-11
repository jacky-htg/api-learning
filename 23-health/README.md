# Rebel Services

Health Check with objective to create  new /health endpoint that returns 200 when the database is ready.

Tasks:
- Create StatusCheck function on libraries/database/database.go
- Add a /health endpoint on routing/route.go
- Create check controller on controllers/checks.go

## File Changes :
- routing/route.go
- libraries/database/database.go

## New File :
- controllers/checks.go

## Adding Dependency :