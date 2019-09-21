# Rebel Services

Token

Create token auth using jwt token.

Task
- add TOKEN_SALT environtment
- create token library on libraries/token/token.go
- add login routing on routing/route.go
- create auth controllers on packages/auth/controllers/auths.go
- create payload login request on packages/auth/payloads/request/login_request.go
- create GetByUsername method on packages/auth/models/user.go
- create payload token response on packages/auth/payloads/response/token_response.go
- create api test for login
 

## File Changes :
- .env
- routing/route.go
- packages/auth/models/user.go
- main_test.go

## New File :
- libraries/token/token.go
- packages/auth/controllers/auths.go
- packages/auth/payloads/request/login_request.go
- packages/auth/payloads/response/token_response.go
- packages/auth/controllers/tests/authstest.go

## Adding Dependency :
- github.com/dgrijalva/jwt-go