# Rebel Services

Authorization using RBAC (role based access controller)  

Tasks:
- design rbac database on schema/migrate.go and schema/seed.go
- run go cmd/main.go migrate && go cmd/main.go seed to dump database
- design rbac routing on routing/route.go
- create scan access command libraries/auth/access.go
- run go cmd/main.go scan-access to insert routing into access table
- create roles service
- update users service to support roles/multi-roles
- create middleware to handle authorization checking 

## File Changes :
- routing/route.go
- models/user.go
- controllers/users.go
- schema/migrate.go
- schema/seed.go
- cmd/main.go

## New File :
- libraries/auth/access.go
- controllers/access.go
- models/access.go
- payloads/response/access_response.go
- controllers/roles.go
- models/roles.go
- payloads/request/roles_request.go
- payloads/response/roles_response.go

## Adding Dependency :