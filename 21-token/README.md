# Rebel Services

Create token auth using jwt token  

Tasks:
- create token library on libraries/token/token.go
- add login routing on routing/route.go
- create auth controllers on controllers/auths.go
- create payload login request on payloads/request/login_request.go
- create GetByUsername method on models/user.go
- create payload token response on payloads/response/token_response.go 

## File Changes :
- routing/route.go
- models/user.go

## New File :
- libraries/token/token.go
- controllers/auths.go
- payloads/request/login_request.go
- payloads/response/token_response.go

## Adding Dependency :
- github.com/dgrijalva/jwt-go