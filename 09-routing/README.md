# Rebel Services

Add a second endpoint: Get.
Add router in routing/route.go to tie the endpoints together.
Update Handler in main.

http://localhost:9000/users
http://localhost:8000/users/1

## File Changes :
Modified main.go
Modified packages/auth/controllers/users.go
Modified packages/auth/models/user.go

## New File :
routing/route.go

## Adding Dependency :
- github.com/go-chi/chi