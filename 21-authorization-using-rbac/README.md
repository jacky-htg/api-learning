# Rebel Services

RBAC

Authorization using RBAC (role based access controller).

Task
- design rbac database on schema/migrate.go and schema/seed.go
- go run cmd/main.go migrate && go run cmd/main.go seed to dump database
- design rbac routing on routing/route.go
- create scan access command libraries/auth/access.go
- go run cmd/main.go scan-access to insert routing into access table
- create roles service
- update users service to support roles/multi-roles
- create middleware to handle authorization checking
- update api testing to support all feature in this chapter
- create api testing for new services (roles & access)

## File Changes :
- routing/route.go
- packages/auth/models/user.go
- packages/auth/controllers/users.go
- schema/migrate.go
- schema/seed.go
- cmd/main.go
- main_test.go
- packages/auth/controllers/tests/userstest.go

## New File :
- libraries/auth/access.go
- packages/auth/controllers/access.go
- packages/auth/models/access.go
- packages/auth/payloads/response/access_response.go
- packages/auth/controllers/roles.go
- packages/auth/models/roles.go
- packages/auth/payloads/request/roles_request.go
- packages/auth/payloads/response/roles_response.go
- middleware/auth.go
- packages/auth/controllers/tests/authstest.go
- packages/auth/controllers/tests/rolestest.go
- packages/auth/controllers/tests/accesstest.go

## Adding Dependency :

## NOTE
In this chapter, we already implementation of handling transaction