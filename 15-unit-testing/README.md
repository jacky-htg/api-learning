# Rebel Services

Unit testing

Task
- create libraries/database/databasetest/docker.go to handle command start and stop mysql container using docker
- create file tests/tests.go to provide main services.
- packages/auth/models/models_test/user_test.go to handle unit test of user models.
- go test packages/auth/models/models_test/user_test.go 

## File Changes :

## New File :
- libraries/database/databasetest/docker.go
- tests/tests.go
- packages/auth/models/models_test/user_test.go

## Adding Dependency :
- github.com/google/go-cmp/cmp