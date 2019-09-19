# Rebel Services

- Add a second endpoint: Get.
- Add router in routing/route.go to tie the endpoints together.
- Create api library for handle app
- Update Handler in main.

## Url
- http://localhost:9000/users
- http://localhost:9000/users/1

## File Changes :
- Modified main.go
- Modified packages/auth/controllers/users.go
- Modified packages/auth/models/user.go

## New File :
- routing/route.go
- libraries/api/app.go

## Adding Dependency :
- github.com/julienschmidt/httprouter